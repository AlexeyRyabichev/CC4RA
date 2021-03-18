package apiinternal

import (
	publicApi "github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/internal/public_api"
	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/api"
	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

var (
	PublicSpecification *api.Specification
	PublicAPI           *api.API
)

func InitializePublic(envName string) error {
	var err error

	logger.Sugar.Info("Parse Public API Specification.")

	if PublicSpecification, err = api.NewSpecification(envName); err != nil {
		return err
	}

	if PublicAPI, err = api.NewAPI(PublicSpecification, publicApi.GetHandler()); err != nil {
		return err
	}

	logger.Sugar.Infof("Start Public API on %s.", PublicSpecification.GetAddr())

	go PublicAPI.Run()

	return nil
}
