/*
 */
package log

import (
	log "github.com/alecthomas/log4go"
	"time"
)

var l log.Logger

// Init 模块初始化调用
func Init(filepath string) {
	var logFilename = filepath
	l = make(log.Logger)

	fileLogWriter := log.NewFileLogWriter(logFilename, false)
	consoleWriter := log.NewConsoleLogWriter()
	consoleWriter.SetFormat("[%T] (%S) %M")

	l.AddFilter("stdout", log.FINEST, consoleWriter)
	l.AddFilter("logfile", log.INFO, fileLogWriter)

	l.Info("Current time is : %s", time.Now().Format("15:04:05 MST 2006/01/02"))

	return
}

// Debug debug level log
func Debug(arg0 interface{}, args ...interface{}) {
	l.Debug(arg0, args)
}

// Info info level log
func Info(arg0 interface{}, args ...interface{}) {
	l.Info(arg0, args)
}

// Warn warn level log
func Warn(arg0 interface{}, args ...interface{}) error {
	return l.Warn(arg0, args)
}

// Error error level log
func Error(arg0 interface{}, args ...interface{}) error {
	return l.Error(arg0, args)
}
