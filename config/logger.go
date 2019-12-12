package config

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func InitLogger() {
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	var output io.Writer
	var err error

	switch Config.Env {
	case EnvStaging, EnvProduction:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		output, err = os.OpenFile(Config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to open log file %v", Config.LogFile)
		}
	case EnvDev:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		output = os.Stdout
	}

	logIO := zerolog.ConsoleWriter{Out: output}

	Logger = zerolog.New(logIO).With().Timestamp().Logger()
}
