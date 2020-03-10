package main

import (
	"github.com/growerlab/svc/config"
	"github.com/growerlab/svc/router"
)

func main() {
	r := router.Router

	router.Route()

	config.Logger.Info().Msgf("booting server on port: %s\n", config.Config.Port)
	if err := r.Run(":" + config.Config.Port); err != nil {
		panic(err)
	}
}
