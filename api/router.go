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
		type Input struct {
			Board    logic.Horizontal `json:"board"`
			Player   uint8/* 1 | 2 */ `json:"player"`
			Position logic.Position `json:"position"`
		}

		input := Input{}

		err := c.BindJSON(&input)
		if err != nil {
			c.Writer.Write([]byte(err.Error()))
			return
		}

		board, wonBy := logic.Add(input.Board, input.Player, input.Position)

		responseJson := struct {
			Board logic.Horizontal `json:"Board"`
			WonBy uint8            `json:"WonBy"`
		}{
			Board: board,
			WonBy: wonBy,
		}

		c.JSON(http.StatusOK, responseJson)
	})

}
