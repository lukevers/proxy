package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func loadAddresses() ([]string, error) {
	data, err := ioutil.ReadFile(*flagAddressesFile)
	if err != nil {
		return nil, err
	}

	addresses := deleteEmpty(strings.Split(string(data), "\n"))
	if *flagDebug {
		fmt.Println(len(addresses), "addresses in allowlist:")
		for _, address := range addresses {
			fmt.Printf("\"%s\"\n", address)
		}
	}

	return addresses, nil
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
