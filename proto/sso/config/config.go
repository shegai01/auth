package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type GRPCConfig struct {
	Port    int    `yaml:"port" env:"GRPC_PORT"`
	Timeout string `yaml:"timeout" env:"GRPC_TIMEOUT"`
}

type Config struct {
	Env            string     `yaml:"env" env-default:"local"`
	DSN            string     `yaml:"POSTGRES_DSN" env-required:"true"`
	MigrationsPath string     `yaml:"migrations_path" env-default:"migrations"`
	GRPC           GRPCConfig `yaml:"grpc"`
}

func MustLoadConfig() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return &cfg
}
