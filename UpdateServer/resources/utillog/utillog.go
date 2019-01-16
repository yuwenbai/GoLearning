package utillog

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	zzxHook "../hook"
)

var _self *UtilLog

//UtilLog 类型定义
type UtilLog struct {
	_logrus *logrus.Logger
}

//Instance 实例
func Instance() *UtilLog {
	if _self == nil {
		_self = new(UtilLog)
		return _self
	}
	return _self
}

//Init 全局维护状态
func (o *UtilLog) Init() {
	_self._logrus = logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		_self._logrus.Out = file
	} else {
		_self._logrus.Info("Failed to log to file, using default stderr")
	}
	_self._logrus.Hooks.Add(zzxHook.NewContextHook())
	_self._logrus.SetLevel(logrus.InfoLevel)
	// _self._logrus.WithFields(logrus.Fields{
	// 	"default": "default",
	// }).Info("初始化log")
}

// //GetLog 获取维护状态
// func (o *UtilLog) GetLog() *logrus.Logger {
// 	return _self._logrus
// }

//Info normal
func (o *UtilLog) Info(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Info(fmt.Sprint(v...))
}

//Error error
func (o *UtilLog) Error(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Error(fmt.Sprint(v...))
}

//Fatal Fatal
func (o *UtilLog) Fatal(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Fatal(fmt.Sprint(v...))
}

//Panic Panic
func (o *UtilLog) Panic(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Panic(fmt.Sprint(v...))
}
