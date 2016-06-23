package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvString(t *testing.T) {
	env := Env("DEBUG")
	assert.Equal(t, "DEBUG", env.String())
	env = Env('d')
	assert.Equal(t, "d", env.String())
}

func TestEnv(t *testing.T) {
	cfg := Config{Port: 80, Environment: "Dev"}
	env, err := cfg.Env()
	if assert.NoError(t, err) {
		assert.Equal(t, Env("Dev"), env)
	}
	cfg = Config{Port: 80, Environment: "Deploy"}
	env, err = cfg.Env()
	if assert.NoError(t, err) {
		assert.Equal(t, Env("Deploy"), env)
	}
	cfg = Config{Port: 80, Environment: "None"}
	env, err = cfg.Env()
	assert.NotNil(t, err)
	assert.Equal(t, Env(""), env)
}
