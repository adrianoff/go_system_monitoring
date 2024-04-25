package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Error(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
}

func New(level string, prefix string) Logger {
	parseLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil
	}
	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(parseLevel),
		DisableCaller:     true,
		Development:       true,
		DisableStacktrace: true,
		Encoding:          "console",
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
	}
	logger := zap.Must(logConfig.Build()).Sugar().Named(prefix)
	return logger
}
