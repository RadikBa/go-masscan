package main

type Operation uint8

const (
	OPERATION_DEFAULT       Operation = iota // nothing specified, so print usage
	OPERATION_List_Adapters                  // --listif
	OPERATION_Selftest                       // --selftest or --regress
	OPERATION_SCAN                           // this is what you expect
	OPERATION_DebugIF                        // --debug if
	OPERATION_ListScan                       // -sL
	OPERATION_ReadScan                       // --readscan <binary-output>
	OPERATION_ReadRange                      // --readrange
	OPERATION_Benchmark                      // --benchmark
	OPERATION_Echo                           // --echo
	OPERATION_EchoAll                        // --echo-all
	OPERATION_EchoCidr                       // --echo-cidr
)

type ScanType struct {
	Tcp    bool
	Udp    bool // -sU
	Sctp   bool
	Ping   bool // --ping, ICMP echo
	Arp    bool // --arp, local ARP scan
	Oproto bool //-sO
}

type NIC struct {
	IfName        string
	Adapter       Adapter
	sourceMac     HardwareAddr
	routerMacIPv4 HardwareAddr
	routerMacIPv6 HardwareAddr
	isVlan        bool
	isUsable      bool
}

type TemplateOptions struct {
}

type Masscan struct {
	ScanType
	op Operation

	echoAll  bool
	topPorts bool
	nics     []NIC

	targets MassIP
	exclude MassIP

	bannerTypes RangeList[uint8]

	// packets-per-second -(-rate parameter)
	maxRate float64
	// --retries or --max-retries parameter (they happen a few seconds apart)
	retries uint32

	isPfRing            bool // --pfring
	isSendQ             bool // --sendq
	isBanners           bool // --banners
	isBannersRawUdp     bool // --rawudp
	isOffline           bool // --offline
	isNoReset           bool // --noreset, don't transmit RST
	isGMT               bool // --gmt, all times in GMT
	isCaptureCert       bool // --capture cert
	isCaptureHtml       bool // --capture html
	isInfinite          bool // -infinite
	isReadScan          bool // -readscan, OPERATION_READSCAN
	isHelloSSL          bool // --ssl
	isHelloSMBv1        bool // --smbv1, instead of v1/v2 hello
	isHelloHTTP         bool // --hello=http, use HTTP on all ports
	isScripting         bool // whether scripting is needed
	isCaptureServerName bool // --capture servername

	/** Packet template options, such as whether we should add a TCP MSS
	 * value, or remove it from the packet */
	templOptions TemplateOptions // e.g. --tcpmss

}
