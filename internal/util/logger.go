package util

import "go.uber.org/zap"

var log *zap.Logger

func SetLogger(logger *zap.Logger) {
	log = logger
}

func Log() *zap.Logger {
	return log
}
