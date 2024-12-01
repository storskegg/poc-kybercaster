package udp

import "net"

const (
	DefaultMulticastIPv4 = "239.0.0.0"
	DefaultMulticastPort = "9999"

	udp4 = "udp4"

	// MaxSafeMTU is the maximum "safe" MTU for UDP packets.
	MaxSafeMTU = 576

	// LgIPHeaderSize
	LgIPHeaderSize = 60

	// UDPHeaderSize
	UDPHeaderSize = 8

	// HeaderPaddingSize is based on the IPSec header size, representing datagram encapsulation.
	HeaderPaddingSize = 64

	// MaxPayloadSize is a conservative, maximum UDP payload size. Please see the readme in this package.
	MaxPayloadSize = MaxSafeMTU - LgIPHeaderSize - UDPHeaderSize - HeaderPaddingSize
)

func DefaultMulticastAddress() string {
	return DefaultMulticastIPv4 + ":" + DefaultMulticastPort
}

func NewCaster(address string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr(udp4, address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP(udp4, nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Listen(address string, handler func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr(udp4, address)
	if err != nil {
		panic(err)
	}

	l, err := net.ListenMulticastUDP(udp4, nil, addr)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			buf := make([]byte, MaxSafeMTU)
			n, addr, err := l.ReadFromUDP(buf)
			if err != nil {
				panic(err)
			}

			handler(addr, n, buf)
		}
	}()
}
