package config

import (
	"io"
	"log"
	"os"
)

const (
	logFilePath = "./logs/blog_api.log"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
	writer io.Writer
}

func NewLogger(p string) (*Logger, error) {
	err := os.MkdirAll("./logs", os.ModePerm)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	logger := log.New(multiWriter, p, log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug:  log.New(file, "DEBUG: ", logger.Flags()),
		info:   log.New(file, "INFO: ", logger.Flags()),
		warn:   log.New(file, "WARN: ", logger.Flags()),
		error:  log.New(file, "ERROR: ", logger.Flags()),
		writer: file,
	}, nil
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}
