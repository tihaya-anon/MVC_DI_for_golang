package cmd

import (
	"MVC_DI/config"
	user_router "MVC_DI/router/user"
	user_controller_builder "MVC_DI/section/user/controller/builder"
	"MVC_DI/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func bindController() {
	userEntryController := user_controller_builder.NewUserEntryControllerBuilder().Build()
	user_router.BindUserEntryController(userEntryController)

	userAuthController := user_controller_builder.NewUserAuthControllerBuilder().Build()
	user_router.BindUserAuthController(userAuthController)
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
	config.InitConfig()
	bindController()
	startServer("/api/v1/public", "/api/v1", 5*time.Second)
}

func Stop() {
	fmt.Println("============= STOP =============")
}
