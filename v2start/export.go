package v2start

import (
	core "github.com/make-money-fast/v2ray-core-v5"
	_ "github.com/make-money-fast/v2ray-core-v5/main/distro/all"
)

func Start(uri string) (core.Server, error) {
	config, err := core.LoadConfig("json", uri)
	if err != nil {
		return nil, err
	}

	server, err := core.New(config)
	if err != nil {
		return nil, err
	}

	return server, nil
}
