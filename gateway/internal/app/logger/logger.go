package logger

import (
	"fmt"
	"os"

	"github.com/sergeyiksanov/golang_project/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(cfg *config.LoggerConfig) (*zap.Logger, error) {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	switch cfg.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "text":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		return nil, fmt.Errorf("invalid log format: %w", cfg.Format)
	}

	cores := make([]zapcore.Core, 0, len(cfg.OutputsPaths))
	for _, output := range cfg.OutputsPaths {
		writer, err := getLogWriter(output)
		if err != nil {
			return nil, err
		}
		core := zapcore.NewCore(encoder, writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	opts := []zap.Option{
		zap.AddCaller(),
	}
	if cfg.EnableStacktrace {
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	}

	return zap.New(combinedCore, opts...), nil
}

func getLogWriter(output string) (zapcore.WriteSyncer, error) {
	switch output {
	case "stdout":
		return zapcore.AddSync(os.Stdout), nil
	case "stderr":
		return zapcore.AddSync(os.Stderr), nil
	default:
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   output,
			MaxSize:    100, //mb
			MaxBackups: 3,
			MaxAge:     30, //days
		}), nil
	}
}
