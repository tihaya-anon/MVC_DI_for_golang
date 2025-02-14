package cmd

import (
	"MVC_DI/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func bindController() {
}

func startServer(publicPath, authPath string, timeOut time.Duration) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := server.NewServer()
	server.Setup(publicPath, authPath)
	server.Run()

	<-ctx.Done()

	server.Stop(timeOut)
}

func Start() {
	bindController()
	startServer("/api/v1/public", "/api/v1", 5*time.Second)
}

func Stop() {
	fmt.Println("============= STOP =============")
}
