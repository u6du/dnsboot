package dnsboot

import (
	"testing"

	"github.com/u6du/dns"
)

func TestRoot(t *testing.T) {
	BootLi(4, &dns.V4)
	BootLi(6, &dns.V6)

}
