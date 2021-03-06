package logger

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"taskbuilder/internal/config"
)

type LogInfo interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type LogFormat interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
}

type LogInfoFormat interface {
	LogFormat
	LogInfo
}

type Logger struct {
	zapLogger *zap.SugaredLogger
}

func (l Logger) Debugf(template string, args ...interface{}) {
	l.zapLogger.Debugf(template, args)
}

func (l Logger) Infof(template string, args ...interface{}) {
	l.zapLogger.Infof(template, args)
}

func (l Logger) Warnf(template string, args ...interface{}) {
	l.zapLogger.Warnf(template, args)
}

func (l Logger) Errorf(template string, args ...interface{}) {
	l.zapLogger.Errorf(template, args)
}

func (l Logger) Fatalf(template string, args ...interface{}) {
	l.zapLogger.Fatalf(template, args)
}

func (l Logger) Panicf(template string, args ...interface{}) {
	l.zapLogger.Panicf(template, args)
}

func (l Logger) Debug(args ...interface{}) {
	l.zapLogger.Debug(args)
}

func (l Logger) Info(args ...interface{}) {
	l.zapLogger.Info(args)
}

func (l Logger) Warn(args ...interface{}) {
	l.zapLogger.Warn(args)
}

func (l Logger) Error(args ...interface{}) {
	l.zapLogger.Error(args)
}

func (l Logger) Fatal(args ...interface{}) {
	l.zapLogger.Fatal(args)
}

func (l Logger) Panic(args ...interface{}) {
	l.zapLogger.Panic(args)
}

func NewLogger(c *config.Config) (LogInfoFormat, error) {
	if c.Logger.Use == "zapLogger" {
		z, err := NewZapLogger(c)
		if err != nil {
			log.Fatalf("failed to initialize ZapLogger %v", err)
			return nil, err
		}
		return &Logger{zapLogger: z}, nil
	}
	return nil, errors.New(fmt.Sprintf("logger not implemented: %s", c.Logger.Use))
}
