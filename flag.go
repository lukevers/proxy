package main

import (
	"flag"
)

var (
	flagDebug         = flag.Bool("debug", false, "Debug/verbose mode")
	flagAddressesFile = flag.String("addresses", "/etc/proxy.addresses.list", "Addresses to only allow from")
	flagBindAddress   = flag.String("bind", ":8080", "Address to bind")
)

func init() {
	flag.Parse()
}
