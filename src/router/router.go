package router

import (
	"MVC_DI/config"
	"MVC_DI/resp"
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type IRegisterRouterFunc = func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup)

var registerRouterFuncList []IRegisterRouterFunc

func RegisterRouter(fn IRegisterRouterFunc) {
	if fn == nil {
		return
	}
	registerRouterFuncList = append(registerRouterFuncList, fn)
}

func ResponseWrapper(fn func(ctx *gin.Context) *resp.IResponse) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusOK, fn(ctx))
	}
}

// # InitRouter
//
// InitRouter initializes the router and starts the server.
// It will iterate over the registerRouterFuncList and call each function
// to register the routes. Then it will start the server and wait for
// the context to be done.
func InitRouter() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// create the gin router
	router := gin.Default()

	publicRouterGroup := router.Group("/api/v1/public")
	authRouterGroup := router.Group("/api/v1")

	for _, registerRouterFunc := range registerRouterFuncList {
		registerRouterFunc(publicRouterGroup, authRouterGroup)
	}

	server := &http.Server{Addr: config.Application.App.Uri, Handler: router}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Start server failed: %s\n", err)
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdowm := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdowm()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %s\n", err)
		return
	}
	fmt.Println("Server shutdown successfully")
}
