package middleware

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/growerlab/codev-svc/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func CtxRepoMiddleware(c *gin.Context) {
	if c.Request.URL.Path == "/graphql" {
		bodyRaw, err := c.GetRawData()
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, errors.WithStack(err))
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyRaw))
		reqOptions := handler.NewRequestOptions(c.Request)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyRaw))

		if strings.Contains(reqOptions.Query, "__type") {
			c.Next()
			return
		}

		reqRepo, err := getRepo(reqOptions.Variables)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		repo, err := model.OpenRepo(reqRepo.Path, reqRepo.Name)
		if err != nil {
			log.Error().Err(err).Msg("failed to open repo")
		} else {
			c.Request = c.Request.WithContext(context.WithValue(c, "repo", repo))
		}
	}
	c.Next()
}

type repoRequest struct {
	Path string
	Name string
}

func (r *repoRequest) fullPath() string {
	return filepath.Join(r.Path, r.Name)
}

func getRepo(variables map[string]interface{}) (*repoRequest, error) {
	var repo = repoRequest{
		Path: variables["path"].(string),
		Name: variables["name"].(string),
	}

	if len(repo.Path) == 0 || len(repo.Name) == 0 {
		return nil, errors.New("not found repo.Path or repo.Name")
	}

	return &repo, nil
}
