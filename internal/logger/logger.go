package loggerinternal

import (
	"go.uber.org/zap"

	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

func Initialize() error {
	var err error

	if logger.Logger, err = zap.NewProduction(); err != nil {
		return err
	}

	logger.Sugar = logger.Logger.Sugar()

	return nil
}
