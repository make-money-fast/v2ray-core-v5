package encoding

import (
	"github.com/make-money-fast/v2ray-core-v5/common/net"
	"github.com/make-money-fast/v2ray-core-v5/common/protocol"
)

//go:generate go run github.com/v2fly/v2ray-core/v5/common/errors/errorgen

const (
	Version = byte(1)
)

var addrParser = protocol.NewAddressParser(
	protocol.AddressFamilyByte(byte(protocol.AddressTypeIPv4), net.AddressFamilyIPv4),
	protocol.AddressFamilyByte(byte(protocol.AddressTypeDomain), net.AddressFamilyDomain),
	protocol.AddressFamilyByte(byte(protocol.AddressTypeIPv6), net.AddressFamilyIPv6),
	protocol.PortThenAddress(),
)
