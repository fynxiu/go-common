package log

import (
	xtime "github.com/fynxiu/go-common/time"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type levelWriterImpl struct {
	logPath     string
	serviceName string
}

func (lw levelWriterImpl) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (lw levelWriterImpl) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	var (
		f *os.File
	)
	f, err = createLogFile(&lw, level)
	if err != nil {
		panic(err)
	}
	return f.Write(p)
}

func createLogFile(lw *levelWriterImpl, level zerolog.Level) (file *os.File, err error) {
	var (
		today    string
		filePath string
		dirPath  string
	)

	// 获取当前日期
	t := time.Now()
	today = t.Format(xtime.LayoutStandardShort)

	// 生成日志文件的名称
	var fileName string
	switch level {
	case zerolog.ErrorLevel:
		fileName = "errors.log"
	case zerolog.InfoLevel:
		fileName = "info.log"
	case zerolog.DebugLevel:
		fileName = "debug.log"
	case zerolog.WarnLevel:
		fileName = "warn.log"
	}

	// 根据环境不同区分不同目录
	sep := string(os.PathSeparator)
	dirPath = lw.logPath + sep + lw.serviceName + sep + today + sep

	// 创建存放日志文件的目录
	if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return
	}

	// 创建日志文件
	filePath = dirPath + fileName

	if file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm); err != nil {
		return
	}

	return
}
