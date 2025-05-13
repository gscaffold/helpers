package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "enjoy yourself!")
	})

	return router
}
