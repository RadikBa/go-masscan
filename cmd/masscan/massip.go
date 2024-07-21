package main

import (
	"net"

	"golang.org/x/exp/constraints"
)

type RangeItem interface {
	net.IP | constraints.Integer
}

type Range[T RangeItem] struct {
	Begin T
	End   T
}

type RangeList[T RangeItem] struct {
	list   []Range[T]
	sorted bool
}

type MassIP struct {
	ipv4 RangeList[net.IP]
	ipv6 RangeList[net.IP]

	ports RangeList[uint16]
}
