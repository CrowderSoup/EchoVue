package config

import (
	"github.com/koding/multiconfig"
	"go.uber.org/fx"
)

// Config our app config
type Config struct {
	Port        int    `default:"1313"`
	Environment string `default:"development"`
	AssetsDir   string `default:"./app/dist"`
}

// LoadConfig loads and returns our app config
func LoadConfig() *Config {
	var config Config
	m := multiconfig.New()
	m.MustLoad(&config)

	return &config
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(
		LoadConfig,
	),
)
