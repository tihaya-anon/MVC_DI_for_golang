package main

import (
	"MVC_DI/global/log"
)

func main() {
	logger := log.GetLogger()

	logger.Debug("This is a debug message.")
	logger.Info("This is an info message.")
	logger.Warn("This is a warn message.")
	logger.Error("This is an error message.")
}
