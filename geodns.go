package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"os/signal"
	"syscall"
	"strconv"

	"github.com/Ne00n/geodns-go/pkg/config"
	"github.com/Ne00n/geodns-go/pkg/query"
	"github.com/oschwald/geoip2-golang"
	"github.com/miekg/dns"
)

func defaultOptions() {
	ex, err := os.Executable()
	if err != nil {
			panic(err)
	}

	exPath := filepath.Dir(ex)

	config.ConfigLocation = flag.String("c", exPath+"/configs/config.yml", "the location of the configuration file of DNS server")
	config.GeoLiteDBLocation = flag.String("g", "/usr/share/GeoIP/GeoLite2-City.mmdb", "the location of GeoLite2/GeoIP2 city MMDB")
	config.Port = flag.Int("p", 53, "which port to listen")
	config.Debug = flag.Bool("D", false, "enable debug mode to print out more information while running the server")
	config.ListenAddress = flag.String("a", "127.0.0.1", "which address to listen for the request")
}

func Serve(port *int, connType string, address *string) {
	srv := &dns.Server{Addr: *address + ":" + strconv.Itoa(*port), Net: connType}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set %s listener %s\n", connType, err.Error())
	}
}

func main() {
	defaultOptions()
	flag.Parse()

	// Initial configMap
	config.FetchConfigMap(config.ConfigLocation)

	// Register domain
	query.RegisterDomain()

	config.GeoDB, _ = geoip2.Open(*config.GeoLiteDBLocation)

	log.Printf("Starting DNS server...\n")

	go Serve(config.Port, "tcp", config.ListenAddress)
	go Serve(config.Port, "udp", config.ListenAddress)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Fatalf("Signal (%v) received, stopping\n", s)
}
