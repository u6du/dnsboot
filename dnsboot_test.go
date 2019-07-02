package dnsboot

import (
	"testing"

	"github.com/u6du/zerolog/debug"
)

func TestRoot(t *testing.T) {
	debug.Msg("test")
	test := func(network uint8) {
		udpLi := BootLi(network)
		t.Logf("ipv%d %s", network, udpLi)
	}
	test(4)
	test(6)

}
