package dnsboot

import (
	"encoding/binary"
	"net"
)

var UDPAddr = map[uint8]func([]byte) []*net.UDPAddr{
	4: addr(4),
	6: addr(16),
}

func addr(byteLen int) func([]byte) []*net.UDPAddr {
	return func(addrLi []byte) []*net.UDPAddr {
		next := 0

		var UDPAddrLi []*net.UDPAddr

		for i := 0; i < len(addrLi); {
			next = i + byteLen
			ip := net.IP(addrLi[i:next])
			i = next
			next = i + 2
			port := binary.LittleEndian.Uint16(addrLi[i:next])
			i = next
			UDPAddrLi = append(UDPAddrLi, &net.UDPAddr{IP: ip, Port: int(port), Zone: ""})
		}

		return UDPAddrLi
	}
}
