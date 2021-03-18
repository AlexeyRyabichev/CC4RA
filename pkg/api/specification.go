package api

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Port         int           `required:"true"`
	WriteTimeout time.Duration `default:"15s"`
	ReadTimeout  time.Duration `default:"15s"`
	IdleTimeout  time.Duration `default:"60s"`

	envName string
}

func NewSpecification(envName string) (*Specification, error) {
	spec := Specification{envName: envName}
	if err := spec.Parse(envName); err != nil {
		return nil, err
	}

	return &spec, nil
}

func (s *Specification) Parse(envName string) error {
	return envconfig.Process(envName, s)
}

func (s *Specification) GetAddr() string {
	return fmt.Sprintf("0.0.0.0:%d", s.Port)
}
