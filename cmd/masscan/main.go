package main

import (
	"fmt"
	"net"
	"scan"
)

func filter(ips []net.IP, black []net.IP) []net.IP {
	return ips
}

type SomeScanner struct {
}

func (s *SomeScanner) Scan(ip net.IP, port uint16) error {
	return nil
}

func main() {
	dst := net.ParseIP("8.8.8.8")
	var scanner scan.Scanner = &SomeScanner{}
	var err error
	if err = scanner.Scan(dst, 80); err != nil {
		fmt.Println(dst.String(), "closed")
	} else {
		fmt.Println(dst.String(), "open")
	}
}
