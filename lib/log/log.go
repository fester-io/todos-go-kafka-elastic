package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	logger, _ = zap.NewDevelopment(zap.AddCallerSkip(1))
	sugar = logger.Sugar()
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
