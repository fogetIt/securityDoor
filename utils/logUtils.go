package utils

import (
	"github.com/astaxie/beego/logs"
)

func Logger() *logs.BeeLogger {
	logger := logs.NewLogger()
	logger.SetLogger(logs.AdapterConsole)
	logger.EnableFuncCallDepth(true)
	return logger
}