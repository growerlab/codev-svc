package main

import (
	"github.com/growerlab/codev-svc/config"
	"github.com/growerlab/codev-svc/router"
)

func main () {
	r := router.Router

	router.SetRouter()

	config.Logger.Info().Msgf("boot server on port: %s\n", config.Config.Port)
	r.Run(":" + config.Config.Port)
	r.Run(":" + config.Config.Port)
}
