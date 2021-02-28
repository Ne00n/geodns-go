package handler

import (
	"github.com/Ne00n/geodns-go/pkg/config"
	"github.com/Ne00n/geodns-go/pkg/fetch"
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
)

func DNSHandler(fqdn string, questionType uint16, sourceAddress net.IP, IPv4 bool) (rr dns.RR) {
	// Get the subdomain information first
	record, err := config.GeoDB.Country(sourceAddress)
	if err != nil {
		log.Panic(err)
	}
	if *(config.Debug) == true {
		log.Printf("Source IP => %s\n", sourceAddress.String())
		log.Printf("Source Country => %s\n", record.Country.IsoCode)
		log.Printf("handler.DNSHandler handling request %s, question type %s\n", fqdn, dns.TypeToString[questionType])
	}
	var value string
	value = ""

	//Go through each domain in config file
	for domain := range config.ConfigMap {
		if strings.Contains(fqdn, domain) {
			// Split FQDN into domain and subdomain
			subdomain := fetch.FetchSubDomainName(fqdn, domain)
			// Doesn't parse bla.bla.example.com but bla.example.com needs fix
			records := fetch.FetchRR(config.ConfigMap[domain])
			qtype := dns.TypeToString[questionType]
			//Check if type is defined
			if typeMatch, ok := records[qtype]; ok {
				typeMap, ok := typeMatch.(map[interface{}]interface{})
				//If the map is empty for that type = no record break up
				if ok == false { break }
				//If the map is not empty check for subdomain otherwise check for any
				if recordMatch, ok := typeMap[subdomain]; ok {
					value = fetch.FetchDefaultValue(recordMatch)
				} else {
					if recordMatch, ok := typeMap["any"]; ok {
						value = fetch.FetchDefaultValue(recordMatch)
					}
				}
			}
		}
	}
	rr = RrGenerator(questionType, fqdn, value)
	return rr
}
