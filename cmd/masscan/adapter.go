package main

import "github.com/google/gopacket/pcap"

type Adapter struct {
	pcap      *pcap.Handle
	sendQueue *pcap.Queue
}
