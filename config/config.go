package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config -.
type (
	Config struct {
		App `yaml:"app"`
	}

	App struct {
		Name string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
