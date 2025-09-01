package logger

import (
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
)

type Config struct {
	Level  string `json:"level" env:"LOG_LEVEL" default:"info"`
	Format string `json:"format" env:"LOG_FORMAT" default:"console"` // "json" or "console"
}

var (
	logger zerolog.Logger
	once   sync.Once
)

// InitLogger initializes the global logger with the given config.
// It is safe to call multiple times; only the first call has effect.
func InitLogger(cfg Config) error {
	var initErr error
	once.Do(func() {
		// Parse log level
		level, err := zerolog.ParseLevel(strings.ToLower(cfg.Level))
		if err != nil {
			level = zerolog.InfoLevel // fallback
		}

		zerolog.SetGlobalLevel(level)

		// Create logger based on format
		if strings.ToLower(cfg.Format) == "json" {
			logger = zerolog.New(os.Stderr).With().
				Timestamp().
				Caller().
				Logger()
		} else {
			logger = zerolog.New(zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: "15:04:05",
			}).With().
				Timestamp().
				Caller().
				Logger()
		}
	})
	return initErr
}

// InitDefault initializes logger with default settings
func InitDefault() error {
	return InitLogger(Config{
		Level:  "info",
		Format: "console",
	})
}

// Helper methods for common log levels
func Info() *zerolog.Event {
	return logger.Info()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

// GetLogger returns the global logger instance
func GetLogger() *zerolog.Logger {
	return &logger
}

// Business-specific helpers
func HTTPRequest(method, path string, statusCode int) {
	Info().
		Str("method", method).
		Str("path", path).
		Int("status", statusCode).
		Msg("HTTP request")
}

func ServerStartup(addr string) {
	Info().Str("addr", addr).Msg("Server starting")
}

func ServerShutdown() {
	Info().Msg("Server shutting down gracefully")

}
