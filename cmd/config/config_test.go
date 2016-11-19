package config_test

import (
	"testing"

	"github.com/marioharper/commuter/cmd/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig_New(t *testing.T) {
	configFile := "_fixtures/config.json"
	theConfig := config.New(configFile)

	assert.Equal(t, theConfig.File, configFile)
}
