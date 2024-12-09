package v2start

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"start",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				time.Sleep(1 * time.Second)
				proxyTest()
			}()
			if err := Start("/Users/zmj/Desktop/projects/v2ray-core-v5/v2start/config.kcp.json"); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func proxyTest() {
	httpPorxy := "http://localhost:11080"
	u, _ := url.Parse(httpPorxy)

	c := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(u),
		},
		Timeout: 3 * time.Second,
	}
	rsp, err := c.Get("https://www.google.com")
	if err != nil {
		fmt.Println("connect failed")
		return
	}
	if rsp.StatusCode == 200 {
		fmt.Println("proxy ok")
	}
}
