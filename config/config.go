package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvProduction = "production"
	EnvStaging    = "staging"
	EnvDev        = "dev"
)

type AppConfig struct {
	Port         string
	Env          string
	GinMode      string
	LogFile      string
	ErrorLogFile string
}

var Config AppConfig

func InitConfig() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = EnvDev
	}
	Config.Env = env

	if _, err := os.Stat(".env." + env); os.IsNotExist(err) {
		// load .env
		fmt.Printf("load config from .env\n")
		godotenv.Load()
	} else {
		fmt.Printf("load config from .env.%s\n", env)
		godotenv.Load(".env." + env)
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}

	Config.Port = "9000"
	port := os.Getenv("PORT")
	if port != "" {
		Config.Port = port
	}

	logFile := os.Getenv("LOG_FILE")
	if logFile != "" {
		Config.LogFile = logFile
	} else {
		fmt.Printf("please specify a value for LOG_FILE\n")
	}

	errorLogFile := os.Getenv("ERROR_LOG_FILE")
	if errorLogFile != "" {
		Config.ErrorLogFile = errorLogFile
	} else {
		fmt.Printf("please specify a value for ERROR_LOG_FILE\n")
	}

	Config.GinMode = os.Getenv("GIN_MODE")

	fmt.Printf("config: %+v\n", Config)
}
