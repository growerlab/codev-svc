package service

import (
	"github.com/growerlab/svc/config"
	"github.com/growerlab/svc/model"
)

func init() {
	config.InitConfig()
	config.InitLogger()
	model.InitRepoConfig()
}
