package utillog

import (
	"bufio"
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
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

	// logDir := "./log/"
	fileName := "server"
	// baseLogPath := logDir + fileName
	writer, err := rotatelogs.New(
		fileName+"-%Y%m%d",
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		// _self._logrus.Fatal("config local file system logger error. %v", errors.WithStack(err))
		_self._logrus.Fatal(err)
	}
	setNull()
	_self._logrus.SetLevel(logrus.InfoLevel)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})

	_self._logrus.AddHook(lfHook)
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	_self._logrus.Out = file
	// } else {
	// 	_self._logrus.Info("Failed to log to file, using default stderr")
	// }
	// _self._logrus.Hooks.Add(zzxHook.NewContextHook())
	// _self._logrus.SetLevel(logrus.InfoLevel)
}
func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	_self._logrus.SetOutput(writer)
}

// //GetLog 获取维护状态
// func (o *UtilLog) GetLog() *logrus.Logger {
// 	return _self._logrus
// }

//Info normal
func (o *UtilLog) Info(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Info(v...)
}

//Error error
func (o *UtilLog) Error(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Error(v...)
}

//Fatal Fatal
func (o *UtilLog) Fatal(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Fatal(v...)
}

//Panic Panic
func (o *UtilLog) Panic(v ...interface{}) {
	_self._logrus.WithFields(logrus.Fields{
		"default": "default",
	}).Panic(v...)
}
