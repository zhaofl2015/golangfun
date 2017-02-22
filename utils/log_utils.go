package utils

import (
	"github.com/astaxie/beego/logs"
)

var (
	Logger *logs.BeeLogger = logs.NewLogger(10000)
)

func Init() {
	Logger.SetLogger("console", "")
	Logger.SetLevel(logs.LevelDebug)
}
