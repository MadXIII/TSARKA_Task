package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(address string, handlers http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           address, // TODO: specify an address ??????
		Handler:        handlers,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
