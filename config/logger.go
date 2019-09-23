package config

import (
  "os"
  "time"
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
)

var Logger *zerolog.Logger

func InitLogger() {
  zerolog.TimestampFunc = func() time.Time {
    return time.Now().UTC()
  }

  file, err := os.Open(Config.LogFile)
  if err != nil {
    log.Fatal().Err(err)
  }


  logger :=zerolog.New(file).Level(zerolog.InfoLevel)
  Logger = &logger
}
