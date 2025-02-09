package cmd

import (
	"MVC_DI/config"
	user_controller_builder "MVC_DI/section/user/controller/builder"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	user_controller_builder.NewUserAuthControllerBuilder().Build().RegisterRoutes(router)
	user_controller_builder.NewUserMetaControllerBuilder().Build().RegisterRoutes(router)
	router.Run(config.Application.App.Uri)
}
