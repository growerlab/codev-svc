package config

import (
	"github.com/growerlab/codev-svc/model"
)

func init() {
	InitConfig()
	InitLogger()
	model.InitRepoDir(Config.RepoDir)
}
