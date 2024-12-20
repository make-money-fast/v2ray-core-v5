package udp

import (
	"context"
	"io"

	"github.com/make-money-fast/v2ray-core-v5/common"
	"github.com/make-money-fast/v2ray-core-v5/common/buf"
	"github.com/make-money-fast/v2ray-core-v5/common/net"
)

type DispatcherI interface {
	common.Closable
	Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer)
}

var DispatcherConnectionTerminationSignalReceiverMark = "DispatcherConnectionTerminationSignalReceiverMark"

type DispatcherConnectionTerminationSignalReceiver interface {
	io.Closer
}
