package log

import (
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

// 定義 16 進制顏色
var (
	timeColor   = color.HEX("#018025")
	callerColor = color.HEX("#777777")
)

var levelColorMap = map[logrus.Level]*color.RGBStyle{
	logrus.PanicLevel: color.HEXStyle("#d32f2f"),
	logrus.FatalLevel: color.HEXStyle("#d32f2f"),
	logrus.ErrorLevel: color.HEXStyle("#d32f2f"),
	logrus.WarnLevel:  color.HEXStyle("#fbc02d"),
	logrus.InfoLevel:  color.HEXStyle("#2196f3"),
	logrus.DebugLevel: color.HEXStyle("#777777"),
}
var levelNameMap = map[logrus.Level]string{
	logrus.PanicLevel: "PANC",
	logrus.FatalLevel: "FATL",
	logrus.ErrorLevel: "ERRO",
	logrus.WarnLevel:  "WARN",
	logrus.InfoLevel:  "INFO",
	logrus.DebugLevel: "DEBU",
}
