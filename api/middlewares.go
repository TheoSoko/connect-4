package api

import (
	"go-projects/connect-4/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hi, there")
}

func boardInit(c *gin.Context) {
	board := logic.Init()
	c.JSON(http.StatusOK, board)

}

func makeAMove(c *gin.Context) {
	input := logic.Input{}

	err := c.BindJSON(&input)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	board, won, err := logic.Add(input.Board, input.Player, input.Position, 4)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  500,
			Message: err.Error(),
		})
		return
	}
	if won == nil {
		*won = false	
	}
	
	responseJson := struct {
		Board logic.Horizontal `json:"board"`
		Won   *bool            `json:"won"`
	}{
		Board: board,
		Won:   won,
	}

	c.JSON(http.StatusOK, responseJson)

}

func diagonalTest(c *gin.Context) {
	input := logic.Input{}

	err := c.BindJSON(&input)
	if err != nil {
		c.Writer.WriteString(err.Error())
		return
	}

	nw_se := logic.SinisterDiagonalScan(input.Board, input.Player, 4)
	ne_sw := logic.DexterDiagonalScan(input.Board, input.Player, 4)

	if nw_se || ne_sw {
		c.Writer.WriteString("Gagné : oui")
		return
	}

	c.Writer.Write([]byte("Gagné : non"))
}
