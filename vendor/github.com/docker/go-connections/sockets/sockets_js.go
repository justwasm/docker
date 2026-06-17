//go:build js && wasm
// +build js,wasm

package sockets

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

const (
	defaultTimeout = 10 * time.Second
)

var ErrProtocolNotAvailable = errors.New("protocol not available")

func ConfigureTransport(tr *http.Transport, proto, addr string) error {
	if tr.MaxIdleConns == 0 {
		tr.MaxIdleConns = 6
		tr.IdleConnTimeout = 30 * time.Second
	}
	switch proto {
	case "unix":
		tr.DisableCompression = true
		dialer := &net.Dialer{
			Timeout: defaultTimeout,
		}
		tr.DialContext = func(ctx context.Context, _, _ string) (net.Conn, error) {
			return dialer.DialContext(ctx, "unix", addr)
		}
	case "npipe":
		return ErrProtocolNotAvailable
	default:
		tr.Proxy = http.ProxyFromEnvironment
		tr.DisableCompression = false
		tr.DialContext = (&net.Dialer{
			Timeout: defaultTimeout,
		}).DialContext
	}
	return nil
}
