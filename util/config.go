package util

import "fmt"

// Env is a type alias for a string
type Env string

func (e Env) String() string {
	return string(e)
}

// Consts for valid environments
const (
	EnvDev    Env = "Dev"
	EnvDeploy Env = "Deploy"
)

// Config struct to hold the port number and environment
type Config struct {
	Port        int    `envconfig:"port" required:"true"`
	Environment string `envconfig:"environment" required:"true"`
}

// Env takes the config and returns a valid Env and an
// error if necessary
func (c Config) Env() (Env, error) {
	switch c.Environment {
	case EnvDev.String():
		return EnvDev, nil
	case EnvDeploy.String():
		return EnvDeploy, nil
	default:
		return Env(""), fmt.Errorf("Invalid Environment")
	}
}
