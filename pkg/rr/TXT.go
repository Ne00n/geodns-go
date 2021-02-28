package rr

import (
	"github.com/miekg/dns"
)

func TypeTXT(fqdn string, value string, ttl uint32) (Rr *dns.TXT) {
	Rr = new(dns.TXT)
	Rr.Hdr = dns.RR_Header{
		Name:   fqdn,
		Rrtype: dns.TypeTXT,
		Class:  dns.ClassINET,
		Ttl:    ttl}
	Rr.Txt[0] = value
	return Rr
}
