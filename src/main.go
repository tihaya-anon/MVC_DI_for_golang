package main

import (
	"MVC_DI/config"
	"fmt"
)

func main() {
	fmt.Printf("init success, your appication config is\n%+v\n", *config.Application)
}
