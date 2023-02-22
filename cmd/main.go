package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	sugar := logger.Sugar()
	sugar.Info("Welcome to FairJudge!")
}
