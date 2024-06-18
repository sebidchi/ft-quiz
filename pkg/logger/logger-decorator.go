package logger

import (
	"github.com/sebidchi/ft-quiz/internal/pkg/domain"
	"go.uber.org/zap"
)

const (
	Warning = 1
	Error   = 2
)

func LogErrors(severity int, err error, logger Logger, items map[string]interface{}) {
	fields := make([]zap.Field, 0)

	for name, value := range items {
		fields = append(fields, zap.Any(name, value))
	}

	if severity == Warning {
		logger.Warn(err.Error(), extraItems(err, fields)...)
	} else {
		logger.Error(err.Error(), extraItems(err, fields)...)
	}

}

func extraItems(error error, fields []zap.Field) []zap.Field {
	if baseError, ok := error.(domain.BaseError); ok {
		for key, value := range baseError.ExtraItems() {
			fields = append(fields, zap.Reflect(key, value))
		}
	}

	fields = append(fields, zap.Error(error))

	return fields
}
