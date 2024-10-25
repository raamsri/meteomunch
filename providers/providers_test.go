package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinkershack/meteomunch/config"
	"github.com/tinkershack/meteomunch/plumber"
)

// Mock implementations
type MockProvider struct{}

func (m *MockProvider) FetchData(coords *plumber.Coordinates) (*plumber.BaseData, error) {
	return &plumber.BaseData{}, nil
}

func (m *MockProvider) SetQueryParams(coords *plumber.Coordinates) {}

func TestNew(t *testing.T) {
	// cfg := &config.Config{}
	cfg, err := config.Get()
	assert.Nil(t, err)

	t.Run("returns open-meteo provider", func(t *testing.T) {
		provider, err := New("open-meteo", cfg)
		assert.NotNil(t, provider)
		assert.Nil(t, err)
	})

	// Mock meteoblue provider to the configuration
	cfg.MeteoProviders = append(cfg.MeteoProviders, config.MeteoProvider{
		Name:    "meteoblue",
		APIKey:  "xxx",
		APIPath: "v1/forecast",
		BaseURI: "https://test-api.meteoblue.com/",
	})

	t.Run("returns meteoblue provider", func(t *testing.T) {
		provider, err := New("meteoblue", cfg)
		assert.NotNil(t, provider)
		assert.Nil(t, err)
	})

	t.Run("returns error for unknown provider", func(t *testing.T) {
		provider, err := New("unknown", cfg)
		assert.Nil(t, provider)
		assert.NotNil(t, err)
		assert.Equal(t, "unknown provider: unknown", err.Error())
	})
}

func TestMockProvider(t *testing.T) {
	provider := &MockProvider{}
	coords := &plumber.Coordinates{}

	t.Run("FetchData returns BaseData", func(t *testing.T) {
		data, err := provider.FetchData(coords)
		assert.NotNil(t, data)
		assert.Nil(t, err)
	})

	t.Run("SetQueryParams does not panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			provider.SetQueryParams(coords)
		})
	})
}
