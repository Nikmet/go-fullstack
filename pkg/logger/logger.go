package logger

import (
	"go-fullstack/config"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(conf *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(conf.Level))
	var logger zerolog.Logger

	if conf.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		cw := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(cw).With().Timestamp().Logger()
	}

	return &logger
}
