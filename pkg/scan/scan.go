package scan

import (
	"net"
)

type Scanner interface {
	Scan(target net.IP, port uint16) error
}

type Option struct {
	Rate        int
	Timeout     int
	FingerPrint bool
	Httpx       bool
}

type OpenIpPort struct {
}
