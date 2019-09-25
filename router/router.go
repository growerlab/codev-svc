package router

import (
	"os"
	"io"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/growerlab/codev-svc/config"
)

var Router *gin.Engine

func init() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	f, err := os.OpenFile(config.Config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		config.Logger.Fatal().Msgf("Failed to create request log file: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(f)

	eF, err := os.OpenFile(config.Config.ErrorLogFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		config.Logger.Fatal().Msgf("Failed to create request error log file: %v", err)
	}
	gin.DefaultErrorWriter = io.MultiWriter(eF)

	Router = gin.New()
	Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	Router.Use(gin.Recovery())

	Router.GET("/ping", func(c *gin.Context) {
    c.String(200, "ok")
  })
}

func SetRouter() {
	Router.POST("/graphql", GraphqlHandler())
	Router.Static("/graphql", "./static")
}
