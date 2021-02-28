package handler

import (
	"github.com/Ne00n/geodns-go/pkg/config"
	rrLib "github.com/Ne00n/geodns-go/pkg/rr"
	"github.com/miekg/dns"
	"log"
)

func RrGenerator(requestType uint16, fqdn string, value string, ttl uint32) (rr dns.RR) {
	switch requestType {
	case dns.TypeA:
		if *(config.Debug) == true {
			log.Printf("Generate A record for %s\n", value)
		}
		rr = rrLib.TypeA(fqdn, value, ttl)
		break
	case dns.TypeAAAA:
		if *(config.Debug) == true {
			log.Printf("Generate AAAA record for %s\n", value)
		}
		rr = rrLib.TypeAAAA(fqdn, value, ttl)
		break
	case dns.TypeCNAME:
		if *(config.Debug) == true {
			log.Printf("Generate CNAME record for %s\n", value)
		}
		rr = rrLib.TypeCNAME(fqdn, value, ttl)
		break
	case dns.TypeTXT:
		if *(config.Debug) == true {
			log.Printf("Generate TXT record for %s\n", value)
		}
		rr = rrLib.TypeTXT(fqdn, value, ttl)
		break
	default:
		rr = new(dns.NULL)
		break
	}
	return rr
}
