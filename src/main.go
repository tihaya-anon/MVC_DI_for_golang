package main

import (
	"MVC_DI/config"
	"MVC_DI/test"
)

func main() {
	config.InitConfig()
	test.Test_JWT()
}
