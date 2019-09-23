package config

import (
  "github.com/joho/godotenv"
  "os"
)

type AppConfig struct {
  Port        string
  LogFile     string
  Env         string
}

var Config AppConfig

func InitConfig() {
  env := os.Getenv("APP_ENV")
  if "" == env {
    env = "dev"
  }
  Config.Env = env

  godotenv.Load(".env." + env + ".local")
  if "test" != env {
    godotenv.Load(".env.local")
  }
  godotenv.Load(".env." + env)
  godotenv.Load()

  Config.Port = "8080"
  port := os.Getenv("PORT")
  if port != "" {
  } else {
    Config.Port = port
  }

  Config.LogFile = "logs/" + Config.Env + ".log"
  logFile := os.Getenv("LOG_FILE")
  if logFile != "" {
      Config.LogFile = logFile
  }
}
