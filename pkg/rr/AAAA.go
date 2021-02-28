package rr

import (
	"github.com/miekg/dns"
	"net"
)

func TypeAAAA(fqdn string, value string, ttl uint32) (Rr *dns.AAAA) {
	Rr = new(dns.AAAA)
	Rr.Hdr = dns.RR_Header{
		Name:   fqdn,
		Rrtype: dns.TypeAAAA,
		Class:  dns.ClassINET,
		Ttl:    ttl}
	Rr.AAAA = net.ParseIP(value)
	return Rr
}
