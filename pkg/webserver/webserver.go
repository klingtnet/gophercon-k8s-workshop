package webserver

import (
	"net"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

type WebServer struct {
	http.Server
}

func New(host, port string, h http.Handler) *WebServer {
	var ws WebServer

	ws.Addr = net.JoinHostPort(host, port)
	ws.Handler = h

	return &ws
}

func (s *WebServer) Start() error {
	return gracehttp.Serve(&s.Server)
}
