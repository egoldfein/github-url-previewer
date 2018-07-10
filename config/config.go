package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

// LoadConfig returns the server config loaded using ENVs or a local env file
func LoadConfig() (*Envs, error) {

	// Get environment.env file
	_, filename, _, _ := runtime.Caller(1)
	envFile := path.Join(path.Dir(filename), "./environment.env")

	if len(envFile) > 0 {
		if err := godotenv.Load(envFile); err != nil {
			return nil, fmt.Errorf("Failure loading envFile: %v", err)
		}
	}

	// Load env vars
	cfg, err := LoadEnvs()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
