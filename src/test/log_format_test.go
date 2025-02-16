package test

import (
	"MVC_DI/global/log"
	"testing"
)

func Test_LogFormat(t *testing.T) {
	logger := log.GetLogger()
	logger.Debug("This is a debug message.")
	logger.Info("This is an info message.")
	logger.Warn("This is a warn message.")
	logger.Error("This is an error message.")
}
