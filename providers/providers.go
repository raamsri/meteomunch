// Package providers offers an interface and a factory method for weather data providers.
// This package is designed to facilitate the integration of various weather data providers
// by defining a common interface that each provider must implement. It also includes a
// factory method to instantiate the appropriate provider based on a given name.
package providers

import (
	"fmt"

	"github.com/tinkershack/meteomunch/config"
	"github.com/tinkershack/meteomunch/logger"
	"github.com/tinkershack/meteomunch/plumber"
)

var l logger.Logger

func init() {
	l = logger.NewTag("providers")
}

// Provider interface defines the methods that each provider must implement
type Provider interface {
	FetchData(coords *plumber.Coordinates) (*plumber.BaseData, error)
	SetQueryParams(coords *plumber.Coordinates)
}

// New returns the appropriate provider based on the name
func New(name string, cfg *config.Config) (Provider, error) {
	switch name {
	case "open-meteo":
		p, err := newOpenMeteo(cfg)
		if err != nil {
			return nil, err
		}
		return p, nil
	case "meteoblue":
		p, err := newMeteoBlue(cfg)
		if err != nil {
			return nil, err
		}
		return p, nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", name)
	}
}
