package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-projects/connect-4/logic"
)

func InitRoutes(router *gin.Engine) {

	router.GET("/", hello)

	router.GET("/lets-go", boardInit)

	router.POST("/play", makeAMove)

	/*
		Tests
	*/
	router.POST("/diagonal-test", func(c *gin.Context) {
		input := logic.Input{}

		err := c.BindJSON(&input)
		if err != nil {
			c.Writer.Write([]byte(err.Error()))
			return
		}

		boule := logic.DexterDiagonalScan(input.Board, input.Player, 4)
		c.Status(http.StatusAccepted)

		if boule {
			c.Writer.Write([]byte("Gagné : oui"))
			return
		}
		c.Writer.Write([]byte("Gagné : non"))

	})

}
