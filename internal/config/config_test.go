package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_GetConfig(t *testing.T) {
	t.Run("get config", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			var cfg Config
			t.Setenv("SERVER_ADDRESS", "localhost:8081")
			err := cfg.GetConfig()
			assert.NoError(t, err)
			assert.Equal(t, Config{Server: "localhost:8081"}, cfg)
		})
		t.Run("server empty", func(t *testing.T) {
			var cfg Config
			err := cfg.GetConfig()
			assert.ErrorIs(t, err, errServerEmpty)
		})
	})

}