package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gscaffold/helpers/app"
	"github.com/gscaffold/helpers/rest"
)

// test curl: curl localhost:8080/v1/1

func main() {
	appIns := app.New()

	restBundle := rest.New(getGinRouter())
	appIns.AddBundle(restBundle)

	appIns.Run(context.Background())
}

func getGinRouter() http.Handler {
	router := rest.NewGinRouter()

	v1 := router.Group("/v1")
	{
		v1.GET("/:id", func(c *gin.Context) {
			id := &rest.IDRequest{}
			if err := c.ShouldBindUri(id); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			switch id.ID {
			case 1:
				c.String(http.StatusOK, "落霞与孤鹜齐飞, 秋水共长天一色.")
				return
			}

			c.String(http.StatusOK, "request id:%d", id)
			return
		})
	}

	return router
}
