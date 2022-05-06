package main

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type websocketResource struct {
	url.URL
}

func (r *websocketResource) Await(ctx context.Context) error {
	netDialer := &net.Dialer{}
	dial := func(network, address string) (net.Conn, error) {
		return netDialer.DialContext(ctx, network, address)
	}
	var timeout time.Duration
	if deadline, ok := ctx.Deadline(); ok {
		timeout = deadline.Sub(time.Now())
	}
	wsDialer := &websocket.Dialer{
		NetDial:          dial,
		HandshakeTimeout: timeout,
		Proxy:            http.ProxyFromEnvironment,
	}

	// IDEA(uwe): Use fragment to specify origin
	// IDEA(uwe): Use fragment to specify subprotocol
	// IDEA(uwe): Use fragment to specify cookies

	conn, _, err := wsDialer.Dial(r.URL.String(), nil)
	if err != nil {
		return &unavailabilityError{err}
	}
	defer func() { _ = conn.Close() }()

	return nil
}
