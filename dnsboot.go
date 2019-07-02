package dnsboot

import (
	"net"

	"github.com/u6du/config"
	"github.com/u6du/dns"
	"github.com/u6du/zerolog/info"
)

var HostBootDefault = "ip.6du.host"
var HostBootPath = "dns/host/boot/"

func BootLi(network uint8, nameserver *dns.Dns) []*net.UDPAddr {
	networkString := string([]byte{network + 48})
	info.Msg(networkString)
	host := networkString + "." + HostBootDefault
	v4host := config.File.OneLine(HostBootPath+networkString, host)

	var ipLi []byte

	r := nameserver.Txt(v4host, func(txt string) bool {
		t, err := Verify(txt)

		if err == nil {
			ipLi = t
			return true
		}
		info.Err(err).End()

		switch err {
		case ErrTimeout:
			ipLi = t
		}
		return false
	})

	if r == nil {
		timeoutCount := 1
		dns.DotTxt(host, func(txt string) bool {
			info.Msg("dot txt")
			t, err := Verify(txt)

			if err == nil {
				ipLi = t
				return true
			}

			switch err {
			case ErrTimeout:
				ipLi = t
				if timeoutCount > 1 {
					return true
				} else {
					timeoutCount++
				}
			}
			return false
		})
	}

	if len(ipLi) > 0 {
		return UDPAddr[network](ipLi)
	}

	return []*net.UDPAddr{}
}
