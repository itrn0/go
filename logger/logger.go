package logger

import (
	"log"
)

var (
	logger Logger = &DefaultLogger{}
)

type DefaultLogger struct {
}

func (*DefaultLogger) Debugf(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (*DefaultLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (*DefaultLogger) Warnf(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (*DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (*DefaultLogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

func (*DefaultLogger) Panicf(format string, args ...interface{}) {
	log.Panicf(format, args)
}

func (*DefaultLogger) Debug(args ...interface{}) {
	log.Print(args)
}

func (*DefaultLogger) Info(args ...interface{}) {
	log.Print(args)
}

func (*DefaultLogger) Warn(args ...interface{}) {
	log.Print(args)
}

func (*DefaultLogger) Error(args ...interface{}) {
	log.Print(args)
}

func (*DefaultLogger) Fatal(args ...interface{}) {
	log.Fatal(args)
}

func (*DefaultLogger) Panic(args ...interface{}) {
	log.Panic(args)
}

func GetLogger() Logger {
	return logger
}

func SetLogger(l Logger) {
	logger = l
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}