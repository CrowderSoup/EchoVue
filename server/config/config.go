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
	DBConfig    DBConfig
}

// DBConfig Config for the database
type DBConfig struct {
	ConnectionString string `default:"data.db"`
	Dialect          string `default:"sqlite3"`
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
