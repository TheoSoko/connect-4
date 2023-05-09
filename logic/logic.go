package logic

import (
	"fmt"
)

type Vertical [6]int // 1 | 2 | 3
type Horizontal [7]Vertical
type Position [2]int

type Input struct {
	Board    Horizontal `json:"board"`
	Player   int/* 1 | 2 */ `json:"player"`
	Position Position `json:"position"`
}

func Init() Horizontal {
	y := Vertical{0, 0, 0, 0, 0, 0}
	x := Horizontal{y, y, y, y, y, y, y}
	return x
}

func Add(board Horizontal, player int /* 1 | 2 */, pos Position) (Board Horizontal, wonBy int) {

	x := pos[0]          // column
	y := pos[1]          // depth
	board[x][y] = player // 1 or 2

	if checkWin(board, player) {
		return board, player
	}

	return board, 0
}

func checkWin(board Horizontal, player int) bool {
	count := int(0)

	for index, col := range board {
		for depth, value := range col {
			if value != player {
				count = 0
				continue
			}
			count++

			fmt.Print("compte à : ", count, "\n")

			if count >= 4 || horizontalScan(board, player, index, depth) || SinisterDiagonalScan(board, player, 4) {
				//fmt.Print("victoire sur colomne : ", index, "\n", "profondeur : ", depth, "\n")
				return true
			}
		}
	}

	return false
}

func horizontalScan(board Horizontal, player int, colIndex int, depth int) bool {
	countX := int(1)

	for i := colIndex; i < len(board); i++ {
		if countX >= 4 {
			return true
		}
		if board[i][depth] != player {
			return false
		}
		countX++
	}

	return false
}

func SinisterDiagonalScan(board Horizontal, player int, power int) bool {
	count := 0

	col := 0
	depth := power - 1
	lineLen := len(board)
	colLen := len(board[0])

	firstCol := true

	for true {
		col2 := col
		depth2 := depth

		// Diagonal sweep
		for true {
			if depth2 == -1 || col2 == lineLen {
				break
			}
			if board[col2][depth2] != player {
				count = 0
			} else {
				count++
			}
			col2++
			depth2--
			if count == 4 {
				break
			}
		}

		if count == 4 {
			return true
		}

		// Up and up on the first col at first
		if firstCol {
			depth++
			if depth == colLen-1 {
				firstCol = false
			}
			continue
		}

		// Then left to right
		col++

		if col == lineLen {
			return false
		}
	}

	return false
}
