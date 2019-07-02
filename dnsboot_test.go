package dnsboot

import (
	"github.com/u6du/dns"

	"testing"
)

func TestBoot(t *testing.T) {
	t.Log("v4", BootLi(4, &dns.V4))

}
