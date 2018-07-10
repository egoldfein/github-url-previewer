package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

const (
	invalidEnvs = "The following envs are invalid : %s"
)

// LogFormats - Available log formats
var LogFormats = map[string]logrus.Formatter{
	"text": &logrus.TextFormatter{},
	"json": &logrus.JSONFormatter{},
}

// Envs is the struct for declaring all the envs
type Envs struct {
	GithubToken string `env:"GITHUB_TOKEN"`
}

// LoadEnvs loads the envs and checks which ones are mandatory
func LoadEnvs() (*Envs, error) {
	cfg := &Envs{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	if errs := validateEnvs(cfg); len(errs) > 0 {
		return nil, fmt.Errorf(invalidEnvs, strings.Join(errs, ", "))
	}

	return cfg, nil
}

func validateEnvs(cfg *Envs) []string {
	errs := make([]string, 0)

	return errs
}
