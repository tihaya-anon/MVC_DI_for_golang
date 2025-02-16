package log

import (
	"MVC_DI/config"
	"fmt"

	"github.com/sirupsen/logrus"
)

type DevFormatter struct{}

func (f *DevFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	time := timeColor.Sprintf("%s", entry.Time.Format("01-02 15:04:05.000"))
	level := levelColorMap[entry.Level].Sprintf(" [%s] ", levelNameMap[entry.Level])
	message := entry.Message
	caller := callerColor.Sprintf("(%s:%d)", (entry.Caller.File), entry.Caller.Line)

	logLine := fmt.Sprintf("%s%s%s\t%s\n", time, level, caller, message)

	return []byte(logLine), nil
}

type ProdFormatter struct{}

func (f *ProdFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message + "\n"), nil
}

func GetLogger() *logrus.Logger {
	if config.Application.Env == "prod" {
		return getProdLogger()
	}
	return getDevLogger()
}

func getDevLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&DevFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	return logger
}

func getProdLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&ProdFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
