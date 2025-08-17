package server

import (
	"auth/internal/config"
	"auth/internal/handler"
	"auth/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) (*Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("server config cannot be nil")
	}

	handler := handler.NewHandler(cfg)
	if handler == nil {
		return nil, fmt.Errorf("failed to create server handler")
	}
	fmt.Println("server handler successfully create")

	router := routes.SetupRouter(handler)

	return &Server{
		cfg:    cfg,
		router: router,
	}, nil
}

func (s *Server) Stop() error {
	fmt.Println("server is stopped")
	return nil
}

func (s *Server) Serve() error {
	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	fmt.Printf("The server is ready to process the request on %s...\n", address)
	return s.router.Run(address)
}
