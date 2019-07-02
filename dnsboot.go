package main

import (
	"github.com/u6du/config"
	"github.com/u6du/dns"
)

var HostBootDefault = "ip.6du.host"
var HostBootPath = "dns/host/boot/"

func BootLi(net string, dns *dns.Dns) {

	v4host := config.File.OneLine(HostBootPath+net, net+"."+HostBootDefault)

	v4txt := dns.Txt(v4host, func(s string) bool {
		println("ipv", net, "  ", s)
		return true
	})
	println(v4txt)
}

func main() {
	BootLi("4", &dns.V4)
	//	var HostTestTxt = config.File.OneLine("dns/host/test/txt", "g.cn")

}
