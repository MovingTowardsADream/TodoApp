package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	// Create Server
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // 1Mb
		ReadTimeout:    10 * time.Second, // 10 sec
		WriteTimeout:   10 * time.Second, // 10 sec
	}
	// Run server and return err if there is a problem
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	// Shutdown server and return err if there is a problem
	return s.httpServer.Shutdown(ctx)
}
