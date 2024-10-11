package logger

import (
	"go.uber.org/zap"
)

func New(level string) *zap.SugaredLogger {
	var cfg zap.Config
	if level == "debug" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	logger, _ := cfg.Build()
	return logger.Sugar()
}
