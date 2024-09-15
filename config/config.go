package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"path"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"clean"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		PG    `yaml:"postgres"`
		Redis `yaml:"redis"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name"    env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Name string `yaml:"name" env:"HTTP_NAME"`
	}

	// Log -.
	Log struct {
		Level string ` yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax  int    ` yaml:"pool_max" env:"PG_POOL_MAX"`
		Username string `env-required:"true" yaml:"username" env:"POSTGRES_USER"`
		Password string `env-required:"true" yaml:"password" env:"POSTGRES_PASSWORD"`
		Host     string `env-required:"true" yaml:"host" env:"POSTGRES_HOST"`
		Port     string `env-required:"true" yaml:"port" env:"POSTGRES_PORT"`
		Name     string `env-required:"true" yaml:"name" env:"POSTGRES_DB"`
	}

	// Redis -.

	Redis struct {
		Address  string `env-required:"true" yaml:"address" env:"REDIS_URL"`
		Db       int    `env-required:"true" yaml:"db" env:"REDIS_DATABASES"`
		Password string `env-required:"true" yaml:"collection" env:"REDIS_PASSWORD"`
	}
)

func New(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
