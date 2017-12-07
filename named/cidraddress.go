package named

import (
	"fmt"
	"net"
)

type CidrAddress struct {
	ip   net.IP
	mask net.IPMask
}

func NewCidrAddress(ip net.IP, mask net.IPMask) *CidrAddress {
	return &CidrAddress{
		ip:   ip,
		mask: mask,
	}
}

func (c CidrAddress) String() string {
	ones, _ := c.mask.Size()
	return fmt.Sprintf("%s/%d", c.ip, ones)
}
