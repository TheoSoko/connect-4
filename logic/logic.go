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

func Add(board Horizontal, player int /* 1 | 2 */, pos Position, variant int) (Board Horizontal, won *bool, err error) {
	x := pos[0]          // column
	y := pos[1]          // depth
	board[x][y] = player // 1 or 2

	win, err := checkWin(board, player, variant)
	if err != nil {
		return board, nil, err
	}

	return board, &win, nil
}

func checkWin(board Horizontal, player int, variant int) (win bool, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
		}
	}()

	power := variant
	count := int(0)

	for index, col := range board {
		for depth, chip := range col {
			if chip != player {
				count = 0
				continue
			}
			count++

			if count >= power || horizontalScan(board, player, index, depth, variant) {
				//fmt.Print("victoire sur colomne : ", index, "\n", "profondeur : ", depth, "\n")
				return true, nil
			}
		}
	}

	if SinisterDiagonalScan(board, player, variant) || DexterDiagonalScan(board, player, variant) {
		return true, nil
	}

	return false, nil
}

func horizontalScan(board Horizontal, player int, colIndex int, depth int, variant int) bool {
	countX := int(1)
	power := variant

	for i := colIndex; i < len(board); i++ {
		if countX >= power {
			return true
		}
		if board[i][depth] != player {
			return false
		}
		countX++
	}

	return false
}

func SinisterDiagonalScan(board Horizontal, player int, variant int) bool {
	power := variant
	lineLen := len(board)
	colLen := len(board[0])
	col := 0
	depth := power - 1
	firstCol := true

	count := 0

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

func DexterDiagonalScan(board Horizontal, player int, variant int) bool {
	power := variant
	lineLen := len(board)
	colLen := len(board[0])
	col := power - 1
	depth := colLen - 1
	lastCol := false

	count := 0

	for true {
		col2 := col
		depth2 := depth

		// Diagonal sweep
		for true {
			if depth2 == -1 || col2 == -1 {
				fmt.Println("In: (break) col2 : ", col2)
				break
			}
			fmt.Println("In: item content : ", board[col2][depth2])

			if board[col2][depth2] == player {
				count++
			} else {
				count = 0
			}
			col2--
			depth2--
			if count == 4 {
				break
			}
		}

		if count == 4 {
			return true
		}

		if lastCol {
			fmt.Println("Out: lastCol here ")
			depth--
			if depth < power-1 {
				fmt.Println("Out: depth at minimum here")
				break
			}
			continue
		}

		col++
		fmt.Println("Out: col =  ", col)

		if col == (lineLen - 1) {
			lastCol = true
			fmt.Println("Out: lastCol activated here ")
		}
	}

	return false
}
