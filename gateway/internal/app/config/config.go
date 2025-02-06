package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var validate = validator.New()

type Config struct {
	Logger LoggerConfig `mapstructure:"logger"`
	Grpc   GrpcConfig   `mapstructure:"grpc"`
}

func Load() (*Config, error) {
	v := viper.New()

	v.SetDefault("logger.level", "info")
	v.SetDefault("logger.format", "json")
	v.SetDefault("logger.output_paths", []string{"stdout"})
	v.SetDefault("logger.enable_stacktrace", false)
	v.SetDefault("grpc.connection_timeout", 10)

	v.SetConfigName("local")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("config read error: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("config unmarshal error: %w", err)
	}

	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validate error: %w", err)
	}

	return &cfg, nil
}
