package all

import (
	"github.com/make-money-fast/v2ray-core-v5/main/commands/all/api"
	"github.com/make-money-fast/v2ray-core-v5/main/commands/all/tls"
	"github.com/make-money-fast/v2ray-core-v5/main/commands/base"
)

//go:generate go run v2ray.com/core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		cmdLove,
		tls.CmdTLS,
		cmdUUID,
		cmdVerify,

		// documents
		docFormat,
		docMerge,
	)
}
