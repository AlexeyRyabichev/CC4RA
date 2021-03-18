package specificationinternal

import (
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

type AppSpecification struct {
	GracefulShutdownTimeout time.Duration

	envName string
}

func NewAppSpecification(envName string) (*AppSpecification, error) {
	spec := AppSpecification{envName: envName}
	if err := spec.Parse(envName); err != nil {
		return nil, err
	}

	return &spec, nil
}

func (s *AppSpecification) Parse(envName string) error {
	return envconfig.Process(envName, s)
}

var Specification *AppSpecification

func Initialize(envName string) error {
	var err error

	logger.Sugar.Info("Parse App AppSpecification.")

	if Specification, err = NewAppSpecification(envName); err != nil {
		return err
	}

	return nil
}
