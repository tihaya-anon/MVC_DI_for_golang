package log

import (
	"MVC_DI/config"
	"MVC_DI/global/module"
	"MVC_DI/util/stream"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type ConsoleFormatter struct{}

func (f *ConsoleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	time := timeColor.Sprintf("%s", entry.Time.Format("15:04:05.000"))
	level := levelColorMap[entry.Level].Sprintf("[%s]", levelNameMap[entry.Level])
	message := entry.Message
	caller := callerColor.Sprintf("%v:%d", entry.Caller.File, entry.Caller.Line)
	arrow := levelColorMap[entry.Level].Sprintf(">>")
	stack := entry.Data["stack"]

	logLine := fmt.Sprintf("%s %s %s %s %s", time, level, caller, arrow, message)

	if stack != nil {
		stack = levelColorMap[entry.Level].Sprintf("%s", entry.Data["stack"])
		logLine = fmt.Sprintf("%s %s %s %s %s\n%s", time, level, caller, arrow, message, stack)
	}

	return append([]byte(logLine), '\n'), nil
}

type ProdFormatter struct{}

type Json struct {
	Time    int    `json:"time"`
	Level   string `json:"level"`
	Caller  string `json:"caller"`
	Message string `json:"message"`
	Stack   string `json:"stack,omitempty"`
}

type FileWriteHook struct {
	level logrus.Level
}

func (hook *FileWriteHook) Levels() []logrus.Level {
	levels:= stream.NewListStream(logrus.AllLevels).Filter(func(level logrus.Level) bool { return level <= hook.level }).ToList()
	fmt.Printf("levels: %v\n", levels)
	return levels
}

func (hook *FileWriteHook) Fire(entry *logrus.Entry) error {
	file, err := os.OpenFile(path.Join(module.GetRoot(), "log", "log.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	jsonData, err := jsonFormat(entry)
	if err != nil {
		return err
	}
	file.Write(jsonData)
	return nil
}

func jsonFormat(entry *logrus.Entry) ([]byte, error) {
	jsonEntry := Json{
		Time:    int(entry.Time.Unix()),
		Level:   entry.Level.String(),
		Caller:  fmt.Sprintf("%v:%d", entry.Caller.File, entry.Caller.Line),
		Message: entry.Message,
	}
	stack := entry.Data["stack"]
	if stack != nil {
		jsonEntry.Stack = stack.(string)
	}
	jsonData, err := json.Marshal(jsonEntry)
	if err != nil {
		return nil, err
	}
	jsonData = append(jsonData, '\n')
	return jsonData, nil
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
	logger.SetFormatter(&ConsoleFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	logger.AddHook(&StackTraceHook{})
	return logger
}

func getProdLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&ConsoleFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	logger.AddHook(&StackTraceHook{})
	logger.AddHook(&FileWriteHook{logrus.InfoLevel})
	return logger
}
