package named

import (
	"fmt"
	"net"
)

type CidrAddress struct {
	Ip   net.IP
	Mask net.IPMask
}

func NewCidrAddress(ip net.IP, mask net.IPMask) *CidrAddress {
	return &CidrAddress{
		Ip:   ip,
		Mask: mask,
	}
}

func (c *CidrAddress) String() string {
	ones, _ := c.Mask.Size()
	return fmt.Sprintf("%s/%d", c.Ip, ones)
}
