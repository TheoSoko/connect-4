package logic

type Vertical [6]uint8 // 1 | 2 | 3
type Horizontal [7]Vertical
type Position [2]uint8

func Init() Horizontal {
	y := Vertical{0, 0, 0, 0, 0, 0}
	x := Horizontal{y, y, y, y, y, y, y}

	return x
}

func Add(board Horizontal, player uint8 /* 1 | 2 */, position Position) (Board Horizontal, wonBy uint8) {
	pos := Position{3, 6}

	x := pos[0]
	y := pos[1]
	board[x][y] = player

	countY := uint8(0)
	countX := uint8(0)

	win := false

	horizontalScan := func(colIndex int, depth int) {
		countX++
		for i := colIndex; i < len(board); i++ {
			if countX >= 4 {
				return
			}
			if board[i][depth] != player {
				countX = 0
				return
			}
			countX++
		}
	}

	for index, col := range board {
		for depth, value := range col {
			if value != player {
				countY--
				continue
			}
			countY++
			horizontalScan(index, depth)
			if countY >= 4 || countX >= 4 {
				win = true
				break
			}
		}
		if win {
			break
		}
	}

	if win {
		return board, player
	}
	return board, 0
}
