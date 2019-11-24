package router

import (
	"github.com/growerlab/codev-svc/schema"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphQLHandler() gin.HandlerFunc{
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GraphiQLHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
