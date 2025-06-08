package web

import (
	"context"
	"net/http"
)

type HttpServer struct {
	httpServer *http.Server
}

func NewHttpServer(httpServer *http.Server) *HttpServer {
	return &HttpServer{
		httpServer: httpServer,
	}
}

func (s *HttpServer) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
