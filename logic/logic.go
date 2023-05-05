package logic

import (
	"fmt"
)

type Vertical [6]uint8 // 1 | 2 | 3
type Horizontal [7]Vertical
type Position [2]uint8

func Init() Horizontal {
	y := Vertical{0, 0, 0, 0, 0, 0}
	x := Horizontal{y, y, y, y, y, y, y}
	return x
}

func Add(board Horizontal, player uint8 /* 1 | 2 */, pos Position) (Board Horizontal, wonBy uint8) {

	x := pos[0]
	y := pos[1]
	board[x][y] = player

	if checkWin(board, player) {
		return board, player
	}

	return board, 0
}

func checkWin(board Horizontal, player uint8) bool {
	count := uint8(0)

	for index, col := range board {
		for depth, value := range col {
			if value != player {
				if (count > 0) {
					count--
				}
				continue
			}
			count++

			fmt.Print("compte à : ", count, "\n")

			if count >= 4 || horizontalScan(board, player, index, depth) {
				fmt.Print("victoire sur colomne : ", index, "\n", "profondeur : ", depth, "\n")
				return true
			}
		}
	}

	return false
}

func horizontalScan(board Horizontal, player uint8, colIndex int, depth int) bool {
	countX := uint8(1)

	for i := colIndex; i < len(board); i++ {
		if countX >= 4 {
			fmt.Print("compte à 4 sur colomne: ", colIndex, "\n")
			return true
		}
		if board[i][depth] != player {
			return false
		}
		countX++
	}

	return false
}
