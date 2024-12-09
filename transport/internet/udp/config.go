package udp

import (
	"github.com/make-money-fast/v2ray-core-v5/common"
	"github.com/make-money-fast/v2ray-core-v5/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
