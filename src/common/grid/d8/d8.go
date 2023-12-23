package d8

import "github.com/glennhartmann/aoc23/src/common"

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
)

func (dir Direction) String() string {
	switch dir {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	case UpLeft:
		return "up-left"
	case UpRight:
		return "up-right"
	case DownLeft:
		return "down-left"
	case DownRight:
		return "down-right"
	default:
		common.Panicf("invalid direction: %v", dir)
	}
	return "INVALID"
}

func GetNextCell(r, c int, dir Direction) (nr, nc int) {
	switch dir {
	case Up:
		return r - 1, c
	case Down:
		return r + 1, c
	case Left:
		return r, c - 1
	case Right:
		return r, c + 1
	case UpLeft:
		return r - 1, c + 1
	case UpRight:
		return r - 1, c + 1
	case DownLeft:
		return r + 1, c - 1
	case DownRight:
		return r + 1, c + 1
	default:
		common.Panicf("invalid direction: %v", dir)
	}
	return -1, -1
}

func OppositeDir(dir Direction) Direction {
	switch dir {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	case UpLeft:
		return DownRight
	case UpRight:
		return DownLeft
	case DownLeft:
		return UpRight
	case DownRight:
		return UpLeft
	default:
		common.Panicf("invalid direction: %v", dir)
	}
	return Direction(-1)
}

func MustFindInStringGrid(lines []string, char byte) (r, c int) {
	for row := range lines {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] == char {
				return row, col
			}
		}
	}
	common.Panicf("%c not found", char)
	return -1, -1
}

func DirForUDLR(c string) Direction {
	switch c {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	case "UL":
		return UpLeft
	case "UR":
		return UpRight
	case "DL":
		return DownLeft
	case "DR":
		return DownRight
	default:
		common.Panicf("invalid direction: %s", c)
	}
	return Direction(-1)
}
