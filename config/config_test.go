package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultConfig(t *testing.T) {
	cfg := NewDefaultConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, "localhost", cfg.Munch.Server.Hostname)
	assert.Equal(t, "50050", cfg.Munch.Server.Port)
	assert.Equal(t, "info", cfg.Munch.LogLevel)
	assert.Equal(t, "mongo", cfg.Mongo.Name)
	assert.Equal(t, "mongodb://localhost:27017", cfg.Mongo.URI)
	assert.Equal(t, "meteomunch", cfg.Mongo.DBName)
	assert.Equal(t, 0, cfg.Mongo.DBNumber)
	assert.Equal(t, "redis", cfg.DLMRedis.Name)
	assert.Equal(t, "redis://localhost:6379", cfg.DLMRedis.URI)
	assert.Equal(t, 1, cfg.DLMRedis.DBNumber)
	assert.Equal(t, 1, len(cfg.MeteoProviders))
	assert.Equal(t, "open-meteo", cfg.MeteoProviders[0].Name)
	assert.Equal(t, "", cfg.MeteoProviders[0].APIKey)
	assert.Equal(t, "v1/forecast", cfg.MeteoProviders[0].APIPath)
	assert.Equal(t, "https://api.open-meteo.com/", cfg.MeteoProviders[0].BaseURI)
}

func TestGet(t *testing.T) {
	cfg, err := Get()
	assert.NotNil(t, cfg)
	assert.NoError(t, err)
}

func TestLoad(t *testing.T) {
	// Test loading default config
	cfg, err := Load(nil)
	assert.NotNil(t, cfg)
	assert.NoError(t, err)

	// Test loading provided config
	providedConfig := &config{
		Munch: Munch{
			Server: MunchServer{
				Hostname: "testhost",
				Port:     "12345",
			},
			LogLevel: "debug",
		},
	}
	cfg, err = Load(providedConfig)
	assert.NotNil(t, cfg)
	assert.NoError(t, err)
	assert.Equal(t, "testhost", cfg.Munch.Server.Hostname)
	assert.Equal(t, "12345", cfg.Munch.Server.Port)
	assert.Equal(t, "debug", cfg.Munch.LogLevel)
}

// TestValidateCriticalFields tests the validateCriticalFields function primarily for its type assertion, not the exhaustive test of all possible errors
func TestValidateCriticalFields(t *testing.T) {
	cfg := NewDefaultConfig()

	// Test with valid config
	err := validateCriticalFields(cfg, false)
	assert.NoError(t, err)

	// Test with missing critical fields
	cfg.MeteoProviders[0].BaseURI = ""
	cfg.MeteoProviders[0].Name = "meteoblue"
	err = validateCriticalFields(cfg, false)
	assert.Error(t, err)
	criticalErrs, ok := err.(*CriticalErrors)
	assert.True(t, ok, "Expected error to be of type *CriticalErrors")
	// assert.Equal(t, 1, len(criticalErrs.Errors))
	// assert.Equal(t, errMsg, criticalErrs.Errors[0].Error())
	errMsg := "Critical config value missing: BaseURI - meteoblue"
	found := false
	for _, e := range criticalErrs.Errors {
		if e.Error() == errMsg {
			found = true
			break
		}
	}
	assert.True(t, found, "Expected error not found in errors: "+errMsg)
}
