package tagged

import (
	"context"

	"github.com/make-money-fast/v2ray-core-v5/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
