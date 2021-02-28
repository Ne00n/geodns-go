package rr

import (
	"github.com/miekg/dns"
)

func TypeCNAME(fqdn string, value string, ttl uint32) (Rr *dns.CNAME) {
	Rr = new(dns.CNAME)
	Rr.Hdr = dns.RR_Header{
		Name:   fqdn,
		Rrtype: dns.TypeCNAME,
		Class:  dns.ClassINET,
		Ttl:    ttl}
	Rr.Target = value
	return Rr
}
