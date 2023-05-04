package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-projects/connect-4/logic"
)

func InitRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hi, there")
	})

	router.GET("/lets-go", func(c *gin.Context) {
		logic.Init()
	})

}
