package logger

import "go.uber.org/zap"

type Logger interface {
	Error(message string, fields ...zap.Field)
	Fatal(message string, fields ...zap.Field)
	Warn(message string, fields ...zap.Field)
	Info(message string, fields ...zap.Field)
}

type NullLogger struct {
}

func (n NullLogger) Error(_ string, _ ...zap.Field) {
}
func (n NullLogger) Fatal(_ string, _ ...zap.Field) {
}
func (n NullLogger) Warn(_ string, _ ...zap.Field) {
}
func (n NullLogger) Info(_ string, _ ...zap.Field) {
}
