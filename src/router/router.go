package router

import (
	"MVC_DI/config"

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

func InitRouter() {
	router := gin.Default()

	publicRouterGroup := router.Group("/api/v1/public")
	authRouterGroup := router.Group("/api/v1")

	for _, registerRouterFunc := range registerRouterFuncList {
		registerRouterFunc(publicRouterGroup, authRouterGroup)
	}

	router.Run(config.Application.App.Uri)
}
