package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

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
	log.Fatal(http.ListenAndServe(*flagBindAddress, proxy))
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
