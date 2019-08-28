package log

import (
	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog"
	"io"
	"log"
	"os"
)

type output string

func (o output) isFile() bool {
	return "file" == o
}

type Log struct {
	loggers     []zerolog.Logger
	serviceName string
}

var l *Log

func Loggers() *Log {
	return l
}

// InitWithConfigFile init logger with config file
func InitWithConfigFile(file string) *Log {
	_, err := toml.DecodeFile(file, &Conf)
	if err != nil {
		log.Fatalf("[log.InitWithConfigFile] failed to get/parse configuration from file: file=%v\n", file)
	}
	serviceName := Conf.ServiceName
	writers := make([]zerolog.Logger, 2)
	var w io.Writer
	for k, v := range Conf.Writers {
		if typ, ok := writerTypes[v.Type]; ok {
			if typ == _file {
				w = levelWriterImpl{
					logPath:     v.Path,
					serviceName: serviceName,
				}
			} else {
				w = os.Stderr
			}
			writers[k] = zerolog.New(w).With().Timestamp().Logger()
		} else {
			log.Fatalf("[log.InitWithConfigFile] wrong logger type: type=%v\n", v.Type)
		}
	}

	l = &Log{
		loggers:     writers,
		serviceName: serviceName,
	}
	return l
}

func (l *Log) Error(msg string) {
	for _, e := range l.loggers {
		e.Error().Msg(msg)
	}
}

func (l *Log) ErrorF(format string, v ...interface{}) {
	for _, e := range l.loggers {
		e.Error().Msgf(format, v...)
	}
}

func (l *Log) Info(msg string) {
	for _, e := range l.loggers {
		e.Info().Msg(msg)
	}
}

func (l *Log) InfoF(format string, v ...interface{}) {
	for _, e := range l.loggers {
		e.Info().Msgf(format, v...)
	}
}

func (l *Log) Debug(msg string) {
	for _, e := range l.loggers {
		e.Debug().Msg(msg)
	}
}

func (l *Log) DebugF(format string, v ...interface{}) {
	for _, e := range l.loggers {
		e.Debug().Msgf(format, v...)
	}
}

/* wrapper */

func Error(msg string) {
	l.Error(msg)
}

func ErrorF(format string, v ...interface{}) {
	l.ErrorF(format, v...)
}

func Info(msg string) {
	l.Info(msg)
}

func InfoF(format string, v ...interface{}) {
	l.InfoF(format, v...)
}

func Debug(msg string) {
	l.Debug(msg)
}

func DebugF(format string, v ...interface{}) {
	l.DebugF(format, v...)
}
