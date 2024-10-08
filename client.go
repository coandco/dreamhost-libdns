package dreamhost

import (
	"github.com/adamantal/go-dreamhost/api"
	"github.com/libdns/libdns"
)

func (p *Provider) init() error {
	client, err := api.NewClient(p.APIKey, nil)
	if err != nil {
		return err
	}
	p.client = *client
	return nil
}

func recordFromApiDnsRecord(apiDnsRecord api.DNSRecord) libdns.Record {
	var rec libdns.Record
	rec.Type = string(apiDnsRecord.Type)
	rec.Value = apiDnsRecord.Value
	// We need to get the name minus the zone to match what libdns expects
	rec.Name = libdns.RelativeName(apiDnsRecord.Record, apiDnsRecord.Zone)
	return rec
}

func apiDnsRecordInputFromRecord(record libdns.Record, zone string) dme.Record {
	var recordInput api.DNSRecordInput
	recordInput.Type = api.RecordType(record.Type)
	recordInput.Value = record.Value
	// Dreamhost expects the record name to be absolute
	recordInput.Record = libdns.AbsoluteName(record.Name, zone)
	return recordInput
}
