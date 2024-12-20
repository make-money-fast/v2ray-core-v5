//go:build !linux && !freebsd
// +build !linux,!freebsd

package tcp

import (
	"github.com/make-money-fast/v2ray-core-v5/common/net"
	"github.com/make-money-fast/v2ray-core-v5/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
