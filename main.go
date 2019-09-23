package main

import (
	"github.com/growerlab/letsgit-svc/config"
	"github.com/growerlab/letsgit-svc/router"
)

func main () {
	r := router.Router

	router.SetRouter()

	r.Run(":" + config.Config.Port)
}
