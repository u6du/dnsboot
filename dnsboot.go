package main

import (
	"github.com/u6du/config"
	"github.com/u6du/dns"
)

var HostBootDefault = "6du.host"

func main() {
	hostPath := "dns/host/boot/"

	v4host := config.File.OneLine(hostPath+"4", "ip4."+HostBootDefault)

	v4txt := dns.ResolveTxtV4(v4host, func(s string) bool {
		println("ip4  ", s)
		return true
	})

	//	var HostTestTxt = config.File.OneLine("dns/host/test/txt", "g.cn")

	println(v4txt)
}
