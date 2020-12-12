package main

import (
	"flag"
)

var (
	flagDebug         = flag.Bool("debug", false, "Debug/verbose mode")
	flagAddressesFile = flag.String("addresses", "/etc/proxy.addresses.list", "Addresses to only allow from")
)

func init() {
	flag.Parse()
}
