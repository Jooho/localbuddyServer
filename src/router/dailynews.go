package router


import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingResponse(router *gin.RouterGroup){

	router.GET("/ping", func(c *gin.Context) {
	c.String(http.StatusOK, "pong")
	})
}