package config

type GrpcConfig struct {
	ServerAddress     string `mapstructure:"level"`
	ConnectionTimeout int    `mapstructure:"connection_timeout"`
}
