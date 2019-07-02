package dnsboot

import (
	"net"

	"github.com/u6du/zerolog/log"
	"github.com/u6du/config"
	"github.com/u6du/dns"
)

var HostBootDefault = "ip.6du.host"
var HostBootPath = "dns/host/boot/"

func BootLi(network uint8, dns *dns.Dns) []*net.UDPAddr {

	networkString := string([]byte{network + 48})
	v4host := config.File.OneLine(HostBootPath+networkString, networkString+"."+HostBootDefault)

	var ipLi []byte

	r := dns.Txt(v4host, func(txt string) bool {
		t, err := Verify(txt)

		if err == nil {
			ipLi = t
			return true
		}

		log.Info().Err(err).Msg("")

		switch err {
		case ErrTimeout:
			ipLi = t
		}
		return false
	})

	if r != nil {
		println("len ipLi ", len(ipLi))

		if len(ipLi) > 0 {
			return UDPAddr[network](ipLi)
		}
	}

	return []*net.UDPAddr{}
}
