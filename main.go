package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

// FOR RASPBERRY PI:
// env GOOS=linux GOARCH=arm GOARM=5 go build

func main() {
	// Build the server
	proxy := goproxy.NewProxyHttpServer()

	if *flagDebug {
		proxy.Verbose = true
	}

	// Security. Only proxy requests from certain IP addresses. IF it errors out,
	if addresses, err := loadAddresses(); err == nil {
		proxy.OnRequest(goproxy.Not(inIPList(addresses...))).HandleConnect(goproxy.AlwaysReject)
	} else {
		fmt.Println("ERROR LOADING ADDRESSES: ", err.Error())
	}

	// Run the server
	log.Fatal(http.ListenAndServe(":8080", proxy))
}

func inIPList(ips ...string) goproxy.ReqConditionFunc {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) bool {
		reqIP, err := getIP(req)
		fmt.Println(reqIP)

		if err != nil {
			return false
		}

		for _, ip := range ips {
			if ip == reqIP {
				return true
			}
		}

		return false
	}
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}

func loadAddresses() ([]string, error) {
	data, err := ioutil.ReadFile(*flagAddressesFile)
	if err != nil {
		return nil, err
	}

	addresses := delete_empty(strings.Split(string(data), "\n"))
	if *flagDebug {
		fmt.Println(len(addresses), "addresses in allowlist:")
		for _, address := range addresses {
			fmt.Printf("\"%s\"\n", address)
		}
	}

	return addresses, nil
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
