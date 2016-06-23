package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/patrickoconnor-wf/go-mandel-service/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockConfig struct {
	mock.Mock
}

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestProcessConfig(t *testing.T) {
	os.Clearenv()
	_, err := processConfig()
	assert.Error(t, err)
	os.Setenv("MANDELSERVICE_PORT", "321")
	os.Setenv("MANDELSERVICE_ENVIRONMENT", "Dev")
	config, err := processConfig()
	assert.Equal(t, util.Config{Port: 321, Environment: "Dev"}, config)
	assert.NoError(t, err)
}
