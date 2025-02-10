package user_router

import (
	"MVC_DI/router"
  user_controller "MVC_DI/section/user/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindUserEntryController (controller *user_controller.UserEntryController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/user-entry"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/user-entry"))

    publicGroup.GET("/hello", controller.Hello)
    authGroup.GET("/hello", controller.Hello)
  })
}