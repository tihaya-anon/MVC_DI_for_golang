package cmd

import (
	test_router "MVC_DI/router/test"
	test_controller_builder "MVC_DI/section/test/controller/builder"
	"MVC_DI/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func bindController() {
	testAController := test_controller_builder.NewTestAControllerBuilder().Build()
	test_router.BindTestAController(testAController)
}

func startServer(publicPath, authPath string, engine *gin.Engine, timeOut time.Duration) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := server.NewServer()
	server.Setup(publicPath, authPath, engine)
	server.Run()

	<-ctx.Done()

	server.Stop(timeOut)
}

func Start() {
	bindController()
	startServer("/api/v1/public", "/api/v1/auth", gin.Default(), 5*time.Second)
}

func Stop() {
	fmt.Println("============= STOP =============")
}
