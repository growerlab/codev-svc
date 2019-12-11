package config

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var Logger zerolog.Logger

func InitLogger() {
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	file, err := os.OpenFile(Config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to open log file %v", Config.LogFile)
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if "staging" != Config.Env || "production" != Config.Env {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logIO := zerolog.ConsoleWriter{Out: file}
	Logger = zerolog.New(logIO).With().Timestamp().Logger()
}
