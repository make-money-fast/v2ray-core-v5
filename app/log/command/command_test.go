package command_test

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/make-money-fast/v2ray-core-v5"
	"github.com/make-money-fast/v2ray-core-v5/app/dispatcher"
	"github.com/make-money-fast/v2ray-core-v5/app/log"
	. "github.com/make-money-fast/v2ray-core-v5/app/log/command"
	"github.com/make-money-fast/v2ray-core-v5/app/proxyman"
	_ "github.com/make-money-fast/v2ray-core-v5/app/proxyman/inbound"
	_ "github.com/make-money-fast/v2ray-core-v5/app/proxyman/outbound"
	"github.com/make-money-fast/v2ray-core-v5/common"
	"github.com/make-money-fast/v2ray-core-v5/common/serial"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
