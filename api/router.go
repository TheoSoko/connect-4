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
		board := logic.Init()
		c.JSON(http.StatusOK, board)
	})

	router.POST("/play", func(c *gin.Context) {

		input := logic.Input{}

		err := c.BindJSON(&input)
		if err != nil {
			c.Writer.Write([]byte(err.Error()))
			return
		}

		board, wonBy := logic.Add(input.Board, input.Player, input.Position, 4)

		responseJson := struct {
			Board logic.Horizontal `json:"Board"`
			WonBy int              `json:"WonBy"`
		}{
			Board: board,
			WonBy: wonBy,
		}

		c.JSON(http.StatusOK, responseJson)
	})

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
