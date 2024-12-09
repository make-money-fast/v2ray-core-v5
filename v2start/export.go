package v2start

import (
	core "github.com/make-money-fast/v2ray-core-v5"
	"github.com/make-money-fast/v2ray-core-v5/main/commands/base"
	_ "github.com/make-money-fast/v2ray-core-v5/main/distro/all"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func Start(uri string) error {
	server, err := startV2Ray(uri)
	if err != nil {
		return err
	}
	if err := server.Start(); err != nil {
		base.Fatalf("Failed to start: %s", err)
	}
	defer server.Close()

	// Explicitly triggering GC to remove garbage from config loading.
	runtime.GC()

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
	return nil
}

func startV2Ray(uri string) (core.Server, error) {
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
