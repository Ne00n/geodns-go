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
	record, err := config.GeoDB.City(sourceAddress)
	if err != nil {
		log.Panic(err)
	}

	var country, region string
	country = strings.ToLower(record.Country.IsoCode)
	region = ""

	//Use regions for us
	if country == "us" {
		if len(record.Subdivisions) > 0 {
			region = country + "-" + strings.ToLower(record.Subdivisions[0].IsoCode)
		}
	}

	if *(config.Debug) == true {
		log.Printf("Source IP => %s\n", sourceAddress.String())
		log.Printf("Source Country => %s\n", country)
		log.Printf("Source Region => %s\n", region)
		log.Printf("handler.DNSHandler handling request %s, question type %s\n", fqdn, dns.TypeToString[questionType])
	}

	var value string
	var ttl uint32
	value = ""
	ttl = 3600

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
				//If the map is an invalid type or empty break up
				if ok == false { break }
				//If the map is not empty check for subdomain otherwise check for any
				if recordMatch, ok := typeMap[subdomain]; ok {
					value = fetch.FetchValue(config.ConfigMap["regions"],recordMatch,country,region)
					ttl =  uint32(fetch.FetchTtlValue(recordMatch))
				} else if recordMatch, ok := typeMap["any"]; ok {
						value = fetch.FetchValue(config.ConfigMap["regions"],recordMatch,country,region)
						ttl =  uint32(fetch.FetchTtlValue(recordMatch))
				}
			}
		}
	}
	rr = RrGenerator(questionType, fqdn, value, ttl)
	return rr
}
