package pastebin

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler, port string) *Server {
	return &Server{
		httpServer:&http.Server{
			Addr: ":" + port,
			Handler: handler,
		},
	}
}

func(s *Server) Run() error {
	fmt.Printf("server started on %s port", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}