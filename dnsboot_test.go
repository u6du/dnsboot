package dnsboot

import (
	"testing"
)

func TestRoot(t *testing.T) {
	v4, v6 := BootLi46()
	t.Logf("boot node\n v4 %s\nv6 %s", v4, v6)

}
