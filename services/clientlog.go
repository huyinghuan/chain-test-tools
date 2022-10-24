package services

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	goutils "git.imgo.tv/bs/gomod-utils"
)

func getRuntime(skip int) (function, filename string, lineno int) {
	function = "???"
	pc, filename, lineno, ok := runtime.Caller(skip)
	if ok {
		function = runtime.FuncForPC(pc).Name()
	}
	filename = filepath.Base(filename)
	return
}

// 包装一层
type ChainClientLog struct{}

var ChainClientLogInstance = ChainClientLog{}

func (c ChainClientLog) Debug(args ...interface{}) {
	goutils.Log.Debug(args...)
}
func (c ChainClientLog) Info(args ...interface{}) {
	goutils.Log.Info(args...)
}
func (c ChainClientLog) Warn(args ...interface{}) {
	functionName, fileName, lineNo := getRuntime(2)
	goutils.ExceptionMonitorAdd("", functionName, fileName, strconv.Itoa(lineNo), "warn")
	goutils.Log.Warning(args...)
}
func (c ChainClientLog) Error(args ...interface{}) {
	functionName, fileName, lineNo := getRuntime(2)
	goutils.ExceptionMonitorAdd("", functionName, fileName, strconv.Itoa(lineNo), "error")
	goutils.Log.Error(args...)
}
func (c ChainClientLog) Debugf(format string, args ...interface{}) {
	goutils.Log.Debugf(format, args...)
}
func (c ChainClientLog) Infof(format string, args ...interface{}) {
	goutils.Log.Infof(format, args...)
}
func (c ChainClientLog) Warnf(format string, args ...interface{}) {
	functionName, fileName, lineNo := getRuntime(2)
	s := fmt.Sprintf("%s=%s %s=%s:%d %s", "Uuid", "", "Runtime", fileName, lineNo, format)
	goutils.ExceptionMonitorAdd("", functionName, fileName, strconv.Itoa(lineNo), "warn")
	goutils.Log.Warningf(s, args...)
}
func (c ChainClientLog) Errorf(format string, args ...interface{}) {
	functionName, fileName, lineNo := getRuntime(2)
	s := fmt.Sprintf("%s=%s %s=%s:%d %s", "Uuid", "", "Runtime", fileName, lineNo, format)
	goutils.ExceptionMonitorAdd("", functionName, fileName, strconv.Itoa(lineNo), "error")
	goutils.Log.Warningf(s, args...)
}
