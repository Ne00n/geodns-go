package fetch

import (
	"github.com/miekg/dns"
	"strings"
)

var (
	ConfigMap map[string]interface{}
)

func FetchValue(regions interface{}, recordMatch interface{}, country string, region string) (result string) {
	if region != "" || country != "" {
		var mapName string
		mapName = FetchMapName(recordMatch)
		regions := regions.(map[interface{}]interface{})
		//Check if the regionMap does exist
		if pops, ok := regions[mapName]; ok {
			pops := pops.(map[interface{}]interface{})
			//Check if the region exist in regionMap
			if value, ok := pops[region]; ok {
				result = value.(string)
			} else if value, ok := pops[country]; ok {
				result = value.(string)
			}
		}
	}
	if result == "" {
		result = FetchDefaultValue(recordMatch)
	}
	return result
}

func FetchMapName(config interface{}) (value string) {
	return config.(map[interface{}]interface{})["map"].(string)
}

func FetchRR(config interface{}) (rrData map[interface{}]interface{}) {
	return config.(map[interface{}]interface{})["records"].(map[interface{}]interface{})
}

func FetchSubDomainName(fqdn string, domain string) (subDomain string) {
	return strings.Split(strings.Split(fqdn, domain)[0], ".")[0]
}

func FetchRrType(rrData interface{}) (rrType string) {
	return rrData.(map[interface{}]interface{})["type"].(string)
}

func FetchDefaultValue(rrData interface{}) (value string) {
	return rrData.(map[interface{}]interface{})["default"].(string)
}

func FetchTtlValue(rrData interface{}) (value int) {
	return rrData.(map[interface{}]interface{})["ttl"].(int)
}

func FetchDNSType(requestType string) (rrType uint16) {
	switch strings.ToUpper(requestType) {
	case "A":
		return dns.TypeA
	case "AAAA":
		return dns.TypeAAAA
	case "CNAME":
		return dns.TypeCNAME
	case "TXT":
		return dns.TypeTXT
	default:
		return dns.TypeNone
	}
}
