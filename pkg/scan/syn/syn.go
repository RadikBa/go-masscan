package syn

import (
	"net"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

	"github.com/RadikBa/go-scan"
)

type SynScanner struct {
	srcMac, gwMac net.HardwareAddr
	devName       string

	gw, srcIp net.IP
	handle    *pcap.Handle
	opts      gopacket.SerializeOptions
	bufPool   *sync.Pool

	option scan.Option
}

func (s *SynScanner) Scan(target net.IP, port uint16) error {
	return nil
}
