package server

import (
	"context"
	"net/http"
)

// Server определяет структуру сервера.
type Server struct {
	httpServer *http.Server
}

// Run запускает сервер.
func (s *Server) Run(host string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: host,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown отключает сервер.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}