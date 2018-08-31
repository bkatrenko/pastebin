package main

import (
	"log"
	"time"

	"github.com/namsral/flag"
)

var cfg Config

func main() {
	var (
		config string
		bind   string
		fqdn   string
		tls    bool
		cert   string
		key    string
		expiry time.Duration
	)

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&cert, "cert", "", "certificate")
	flag.StringVar(&key, "key", "", "certificate key")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.StringVar(&fqdn, "fqdn", "localhost", "FQDN for public access")
	flag.DurationVar(&expiry, "expiry", 5*time.Minute, "expiry time for pastes")
	flag.BoolVar(&tls, "tls", false, "do you want some security?")

	flag.Parse()

	if expiry.Seconds() < 60 {
		log.Fatalf("expiry of %s is too small", expiry)
	}

	// TODO: Abstract the Config and Handlers better
	cfg.fqdn = fqdn
	cfg.expiry = expiry

	if tls {
		if cert == "" {
			panic("empty certificate key")
		}

		if key == "" {
			panic("empty key")
		}

		NewServer(bind, cfg).ListenAndServeTLS(cert, key)
		return
	}

	NewServer(bind, cfg).ListenAndServe()
}
