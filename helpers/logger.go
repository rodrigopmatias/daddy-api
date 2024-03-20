package helpers

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

var logger *Logger

func GetLogger() *Logger {
	if logger == nil {
		logger = newLogger()
	}

	return logger
}

func newLogger() *Logger {
	writer := io.Writer(os.Stdout)
	opts := log.Ldate | log.Ltime

	return &Logger{
		debug:  log.New(writer, "DEBUG ", opts),
		info:   log.New(writer, "INFO  ", opts),
		warn:   log.New(writer, "WARN  ", opts),
		err:    log.New(writer, "ERROR ", opts),
		writer: writer,
	}
}

func (l Logger) Debug(v ...any) {
	l.debug.Println(v...)
}

func (l Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}

func (l Logger) Info(v ...any) {
	l.info.Println(v...)
}

func (l Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l Logger) Warn(v ...any) {
	l.warn.Println(v...)
}

func (l Logger) Warnf(format string, v ...any) {
	l.warn.Printf(format, v...)
}

func (l Logger) Err(v ...any) {
	l.err.Println(v...)
}

func (l Logger) Errf(format string, v ...any) {
	l.err.Printf(format, v...)
}
