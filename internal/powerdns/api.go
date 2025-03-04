package powerdns

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/k3llla/terraform-provider-powerdns/internal/powerdns/client"
	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/client/zones"
	"github.com/k3llla/terraform-provider-powerdns/internal/powerdns/models"
)

type Client struct {
	client   *apiclient.PowerDNSAuthoritativeHTTPAPI
	authInfo runtime.ClientAuthInfoWriter
}

type Zone struct {
	ID         string
	Name       string
	Kind       string
	DNSSec     bool
	Serial     int64
	Masters    []string
	RecordSets []RecordSet
}

type RecordSet struct {
	Name       string
	Type       string
	TTL        int64
	Changetype string
	Records    []string
}

func New(ctx context.Context, apiKey, serverHost, basePath, scheme string) *Client {
	transport := httptransport.New(serverHost, basePath, []string{scheme})

	return &Client{
		client:   apiclient.New(transport, strfmt.Default),
		authInfo: httptransport.APIKeyAuth("X-API-Key", "header", apiKey),
	}
}

func (pdns *Client) CreateZone(ctx context.Context, serverID string, zone *Zone) (*Zone, error) {
	if zone.Name == "" {
		return nil, errors.New("zone name is required")
	}
	if zone.Kind == "" {
		return nil, errors.New("zone kind is required")
	}
	zoneStruct := transformZoneToAPI(zone)

	responseRrsets := true
	params := zones.NewCreateZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneStruct(zoneStruct).WithRrsets(&responseRrsets)

	resp, err := pdns.client.Zones.CreateZone(params, pdns.authInfo)
	if err != nil {
		return nil, err
	}

	return transformAPIToZone(resp.Payload), nil
}

func (pdns *Client) UpdateZone(ctx context.Context, serverID, zoneID string, zone *Zone) error {
	zoneStruct := transformZoneToAPI(zone)

	params := zones.NewPutZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneID(zoneID).WithZoneStruct(zoneStruct)

	_, err := pdns.client.Zones.PutZone(params, pdns.authInfo)
	return err
}

func (pdns *Client) GetZone(ctx context.Context, serverID, zoneID string) (*Zone, error) {
	params := zones.NewListZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneID(zoneID)

	resp, err := pdns.client.Zones.ListZone(params, pdns.authInfo)
	if err != nil {
		return nil, err
	}

	return transformAPIToZone(resp.Payload), nil
}

func (pdns *Client) DeleteZone(ctx context.Context, serverID, zoneID string) error {
	params := zones.NewDeleteZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneID(zoneID)

	_, err := pdns.client.Zones.DeleteZone(params, pdns.authInfo)
	return err
}

func (pdns *Client) CreateRecordSet(ctx context.Context, serverID, zoneID string, recordSet *RecordSet) (*RecordSet, error) {
	rrset := transformRecordSetToAPI(recordSet)

	changeTypeReplace := "REPLACE"
	rrset.Changetype = &changeTypeReplace

	zone := &models.Zone{
		Rrsets: []*models.RRSet{rrset},
	}

	params := zones.NewPatchZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneID(zoneID).WithZoneStruct(zone)

	if _, err := pdns.client.Zones.PatchZone(params, pdns.authInfo); err != nil {
		return nil, err
	}

	return pdns.GetRecordSet(ctx, serverID, zoneID, recordSet.Name, recordSet.Type)
}

func (pdns *Client) UpdateRecordSet(ctx context.Context, serverID, zoneID string, recordSet *RecordSet) error {
	rrset := transformRecordSetToAPI(recordSet)

	changeTypeReplace := "REPLACE"
	rrset.Changetype = &changeTypeReplace

	zone := &models.Zone{
		Rrsets: []*models.RRSet{rrset},
	}

	params := zones.NewPatchZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneID(zoneID).WithZoneStruct(zone)

	_, err := pdns.client.Zones.PatchZone(params, pdns.authInfo)
	return err
}

func (pdns *Client) GetRecordSet(ctx context.Context, serverID, zoneID, recordSetName, recordSetType string) (*RecordSet, error) {
	params := zones.NewListZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneID(zoneID)

	resp, err := pdns.client.Zones.ListZone(params, pdns.authInfo)
	if err != nil {
		return nil, err
	}

	rrs := []*models.RRSet{}
	for _, rrset := range resp.Payload.Rrsets {
		if *rrset.Name == recordSetName {
			rrs = append(rrs, rrset)
		}
	}

	switch len(rrs) {
	case 0:
		return nil, fmt.Errorf("record set '%s' not found", recordSetName)
	case 1:
		return transformAPIToRecordSet(rrs[0]), nil
	default:
		if recordSetType == "" {
			return nil, errors.New("multiple record sets found with the same name, type required")
		}

		for _, rrset := range rrs {
			if *rrset.Type == recordSetType {
				return transformAPIToRecordSet(rrset), nil
			}
		}
		return nil, fmt.Errorf("record set '%s' with type '%s' not found", recordSetName, recordSetType)
	}
}

func (pdns *Client) DeleteRecordSet(ctx context.Context, serverID, zoneID string, recordSet *RecordSet) error {
	rrset := transformRecordSetToAPI(recordSet)

	changeTypeDelete := "DELETE"
	rrset.Changetype = &changeTypeDelete

	zone := &models.Zone{
		Rrsets: []*models.RRSet{rrset},
	}

	params := zones.NewPatchZoneParamsWithContext(ctx).WithServerID(serverID).WithZoneID(zoneID).WithZoneStruct(zone)

	_, err := pdns.client.Zones.PatchZone(params, pdns.authInfo)
	return err
}

func transformRecordSetToAPI(recordSet *RecordSet) *models.RRSet {
	records := make([]*models.Record, len(recordSet.Records))
	for i, record := range recordSet.Records {
		record := record
		records[i] = &models.Record{
			Content:  &record,
			Disabled: false,
		}
	}
	return &models.RRSet{
		Name:    &recordSet.Name,
		Type:    &recordSet.Type,
		TTL:     &recordSet.TTL,
		Records: records,
	}
}

func transformAPIToRecordSet(rrset *models.RRSet) *RecordSet {
	records := make([]string, len(rrset.Records))
	for i, record := range rrset.Records {
		records[i] = *record.Content
	}
	return &RecordSet{
		Name:    *rrset.Name,
		Type:    *rrset.Type,
		TTL:     *rrset.TTL,
		Records: records,
	}
}

func transformZoneToAPI(zone *Zone) *models.Zone {
	rrsets := make([]*models.RRSet, len(zone.RecordSets))
	for i, recordset := range zone.RecordSets {
		records := make([]*models.Record, len(recordset.Records))
		for j, record := range recordset.Records {
			records[j] = &models.Record{
				Content: &record,
			}
		}
		rrsets[i] = &models.RRSet{
			Name:    &recordset.Name,
			Type:    &recordset.Type,
			TTL:     &recordset.TTL,
			Records: records,
		}
	}

	return &models.Zone{
		Name:    zone.Name,
		Kind:    zone.Kind,
		Dnssec:  zone.DNSSec,
		Serial:  zone.Serial,
		Masters: zone.Masters,
		Rrsets:  rrsets,
	}
}

func transformAPIToZone(zone *models.Zone) *Zone {
	recordsets := make([]RecordSet, len(zone.Rrsets))
	for i, rrset := range zone.Rrsets {
		records := make([]string, len(rrset.Records))
		for j, record := range rrset.Records {
			records[j] = *record.Content
		}
		recordsets[i] = RecordSet{
			Name:    *rrset.Name,
			Type:    *rrset.Type,
			TTL:     *rrset.TTL,
			Records: records,
		}
	}

	return &Zone{
		ID:         zone.ID,
		Name:       zone.Name,
		Kind:       zone.Kind,
		DNSSec:     zone.Dnssec,
		Serial:     zone.Serial,
		Masters:    zone.Masters,
		RecordSets: recordsets,
	}
}
