package cmd

import (
	"MVC_DI/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.Run(config.Application.App.Uri)
}
