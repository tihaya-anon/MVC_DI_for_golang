package server

import (
	"MVC_DI/config"
	"MVC_DI/router"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Setup(publicPath, authPath string) {
	engine := gin.Default()

	publicRouterGroup := engine.Group(publicPath)
	authRouterGroup := engine.Group(authPath)

	for _, registerRouterFunc := range router.RegisterRouterFuncList {
		registerRouterFunc(publicRouterGroup, authRouterGroup)
	}

	s.server = &http.Server{
		Addr:    config.Application.App.Uri,
		Handler: engine,
	}
}

func (s *Server) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Start server failed: %s\n", err)
			return
		}
	}()
}

func (s *Server) Stop(timeOut time.Duration) {
	ctx, cancelShutdowm := context.WithTimeout(context.Background(), timeOut)
	defer cancelShutdowm()

	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %s\n", err)
		return
	}
	fmt.Println("Server shutdown successfully")
}
