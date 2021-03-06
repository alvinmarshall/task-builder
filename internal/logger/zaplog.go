package logger

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"taskbuilder/internal/config"
)

func NewZapLogger(config *config.Config) (*zap.SugaredLogger, error) {
	zapCfg := zap.Config{}
	switch strings.ToLower(config.Logger.Environment) {
	case "dev", "development", "local":
		zapCfg = zap.NewDevelopmentConfig()
	case "prod", "production":
		zapCfg = zap.NewProductionConfig()
	default:
		return nil, errors.New("logger environment not implemented")
	}
	level := setLevel(config.Logger.LogLevel)
	zapCfg.Level = zap.NewAtomicLevelAt(level)
	zapCfg.OutputPaths = []string{config.Logger.FileName}
	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

func setLevel(level string) zapcore.Level {
	var zapLevel zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "warn", "warning":
		zapLevel = zap.WarnLevel
	}
	return zapLevel
}
