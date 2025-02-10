package cmd

import (
	"MVC_DI/config"
	"MVC_DI/router"
	"fmt"
)

func Start() {
	config.InitConfig()
	router.InitRouter()
}

func Stop() {
	fmt.Println("============= STOP =============")
}
