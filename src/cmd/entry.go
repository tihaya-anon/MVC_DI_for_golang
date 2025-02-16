package cmd

import (
	"MVC_DI/config"
	"MVC_DI/global/log"
	test_router "MVC_DI/router/test"
	test_controller_builder "MVC_DI/section/test/controller/builder"
	test_mapper_builder "MVC_DI/section/test/mapper/builder"
	test_service_builder "MVC_DI/section/test/service/builder"
	"MVC_DI/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ControllerProxy struct{}

func (c *ControllerProxy) Before(ctx *gin.Context, logger *logrus.Logger) {
	logger.Info("In controller, before service")
}

func (c *ControllerProxy) After(ctx *gin.Context, logger *logrus.Logger) {
	logger.Info("In controller, after service")
}

type ServiceProxy struct{}

func (c *ServiceProxy) Before(ctx *gin.Context, logger *logrus.Logger) {
	logger.Info("In service, before mapper")
}

func (c *ServiceProxy) After(ctx *gin.Context, logger *logrus.Logger) {
	logger.Info("In service, after mapper")
}

func bindController() {
	testAMapper := test_mapper_builder.NewTestAMapperBuilder().WithLogger(log.GetLogger()).Build()
	testAService := test_service_builder.NewTestAServiceBuilder().WithLogger(log.GetLogger()).WithProxy(&ServiceProxy{}).WithTestAMapper(testAMapper).Build()
	testAController := test_controller_builder.NewTestAControllerBuilder().WithLogger(log.GetLogger()).WithProxy(&ControllerProxy{}).WithTestAService(testAService).Build()
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
	fmt.Println("============= START =============")
	bindController()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	fmt.Printf("listen to: %s\n", config.Application.App.Uri)
	startServer("/api/v1/public", "/api/v1/auth", engine, 5*time.Second)
}

func Stop() {
	fmt.Println("============= STOP =============")
}
