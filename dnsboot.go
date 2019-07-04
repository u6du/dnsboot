package dnsboot

import (
	"net"
	"sync"

	"github.com/u6du/config"
	"github.com/u6du/dns"
	"github.com/u6du/zerolog/warn"
)

var HostBootDefault = "ip.6du.host"
var HostBootPath = "dns/host/boot/"

func BootLi46() ([]*net.UDPAddr, []*net.UDPAddr) {
	wait := sync.WaitGroup{}
	wait.Add(2)
	var v4, v6 []*net.UDPAddr
	go func() {
		defer wait.Done()
		v4 = BootLi(4)
	}()
	go func() {
		defer wait.Done()
		v6 = BootLi(6)
	}()
	wait.Wait()
	return v4, v6
}

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
						warn.Msg("ipv" + networkString + " boot ip dns txt record expired")
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
