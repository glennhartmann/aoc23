package d4

import (
	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/grid/d8"
)

type Direction int

const (
	Up    = Direction(d8.Up)
	Down  = Direction(d8.Down)
	Left  = Direction(d8.Left)
	Right = Direction(d8.Right)
)

func (dir Direction) String() string {
	return d8.Direction(dir).String()
}

func GetDirChar(dir Direction) byte {
	switch dir {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	default:
		common.Panicf("invalid direction: %v", dir)
	}
	return '!'
}

func GetNextCell(r, c int, dir Direction) (nr, nc int) {
	return d8.GetNextCell(r, c, d8.Direction(dir))
}

func OppositeDir(dir Direction) Direction {
	return Direction(d8.OppositeDir(d8.Direction(dir)))
}

func MustFindInStringGrid(lines []string, c byte) (x, y int) {
	return d8.MustFindInStringGrid(lines, c)
}
