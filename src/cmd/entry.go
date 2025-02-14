package cmd

import (
	"MVC_DI/config"
	"MVC_DI/router"
	user_router "MVC_DI/router/user"
	user_controller_builder "MVC_DI/section/user/controller/builder"
	"fmt"
)

func Start() {
	config.InitConfig()
	userEntryController := user_controller_builder.NewUserEntryControllerBuilder().Build()
	userAuthController := user_controller_builder.NewUserAuthControllerBuilder().Build()
	user_router.BindUserEntryController(userEntryController)
	user_router.BindUserAuthController(userAuthController)
	router.InitRouter()
}

func Stop() {
	fmt.Println("============= STOP =============")
}
