package config

type LoggerConfig struct {
	Level            string   `mapstructure:"level" validate:"oneof=debug info warn error"`
	Format           string   `mapstructure:"format" validate:"oneof=json text"`
	OutputsPaths     []string `mapstructure:"output_paths"`
	EnableStacktrace bool     `mapstructure:"enable_stacktrace"`
}
