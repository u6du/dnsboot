package dnsboot

import (
	"net"

	"github.com/u6du/config"
	"github.com/u6du/dns"
)

var HostBootDefault = "ip.6du.host"
var HostBootPath = "dns/host/boot/"

func BootLi(network uint8) []*net.UDPAddr {
	nameserver := dns.DNS[network]
	networkString := string([]byte{network + 48})
	host := networkString + "." + HostBootDefault
	v4host := config.File.OneLine(HostBootPath+networkString, host)

	var ipLi []byte

	r := nameserver.Txt(v4host, func(txt string) bool {
		t, err := Verify(txt)

		if err == nil {
			ipLi = t
			return true
		}

		switch err {
		case ErrTimeout:
			if len(t) > 0 {
				ipLi = t
			}
		}
		return false
	})

	if r == nil {
		if len(ipLi) > 0 || nameserver.TxtTest(config.File.OneLine("dns/host/test/txt", "g.cn")) {
			timeoutCount := 1
			dns.DotTxt(host, func(txt string) bool {
				t, err := Verify(txt)

				if err == nil {
					ipLi = t
					return true
				}

				switch err {
				case ErrTimeout:
					if len(t) > 0 {
						ipLi = t
					}
					if timeoutCount > 1 {
						return true
					} else {
						timeoutCount++
					}
				}
				return false
			})
		}
	}

	if len(ipLi) > 0 {
		return UDPAddr[network](ipLi)
	}

	return []*net.UDPAddr{}
}
