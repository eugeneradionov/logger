package logger

import (
	"go.uber.org/zap"
)

func Load(preset LogPreset) (*zap.Logger, error) {
	return NewConfig(preset).Build()
}
