package dnsboot

import (
	"testing"
)

func TestRoot(t *testing.T) {

	test := func(network uint8) {
		udpLi := BootLi(network)
		t.Logf("ipv%d %s", network, udpLi)
	}
	test(4)
	test(6)

}
