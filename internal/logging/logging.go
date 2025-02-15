package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer   []io.Writer
	LogLevel []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevel
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(s string, v interface{}) Logger {
	return Logger{l.WithField(s, v)}
}

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err := os.MkdirAll("logs", 0744)
	if err != nil {
		panic(err)
	}
	allFile, err := os.OpenFile("logs/result.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}
	l.SetOutput(io.Discard)
	l.AddHook(&writerHook{
		Writer:   []io.Writer{allFile, os.Stdout},
		LogLevel: logrus.AllLevels,
	})
	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
