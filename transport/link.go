package transport

import "github.com/make-money-fast/v2ray-core-v5/common/buf"

// Link is a utility for connecting between an inbound and an outbound proxy handler.
type Link struct {
	Reader buf.Reader
	Writer buf.Writer
}
