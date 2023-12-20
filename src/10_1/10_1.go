package main

import (
	"log"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/grid/d4"
	"github.com/glennhartmann/aoc23/src/common/must"
)

func main() {
	lines := must.GetFullInput()
	lines = common.AddSentinal(lines, ".")

	x, y := d4.MustFindInStringGrid(lines, 'S')
	s := determineStartShape(lines, x, y)
	log.Printf("start [%d, %d] is %c", x, y, s)

	nodesVisited := 1
	dirs := getDirections(s)
	dir := dirs[0]
	for {
		x, y = d4.GetNextCell(x, y, dir)
		log.Printf("moving %v to [%d, %d]", dir, x, y)

		if lines[y][x] == 'S' {
			break
		}
		nodesVisited++
		dir = getOtherDir(lines[y][x], d4.OppositeDir(dir))
	}
	log.Printf("visited %d nodes", nodesVisited)
	log.Printf("farthest distance: %d", nodesVisited/2)
}

func determineStartShape(lines []string, x, y int) byte {
	cl, cr, cu, cd := connectsRight(lines[y][x-1]), connectsLeft(lines[y][x+1]), connectsDown(lines[y-1][x]), connectsUp(lines[y+1][x])
	log.Printf("start connections: left %v, right %v, up %v, down %v", cl, cr, cu, cd)
	switch {
	case cu && cd:
		return '|'
	case cl && cr:
		return '-'
	case cu && cr:
		return 'L'
	case cu && cl:
		return 'J'
	case cd && cl:
		return '7'
	case cd && cr:
		return 'F'
	default:
		panic("bad start shape")
	}
}

func connectsLeft(p byte) bool {
	return p == '-' || p == 'J' || p == '7'
}

func connectsRight(p byte) bool {
	return p == '-' || p == 'L' || p == 'F'
}

func connectsUp(p byte) bool {
	return p == '|' || p == 'L' || p == 'J'
}

func connectsDown(p byte) bool {
	return p == '|' || p == '7' || p == 'F'
}

func getDirections(p byte) []d4.Direction {
	switch p {
	case '|':
		return []d4.Direction{d4.Up, d4.Down}
	case '-':
		return []d4.Direction{d4.Left, d4.Right}
	case 'L':
		return []d4.Direction{d4.Up, d4.Right}
	case 'J':
		return []d4.Direction{d4.Up, d4.Left}
	case '7':
		return []d4.Direction{d4.Down, d4.Left}
	case 'F':
		return []d4.Direction{d4.Down, d4.Right}
	default:
		panic("bad direction")
	}
}

func getOtherDir(p byte, d d4.Direction) d4.Direction {
	dirs := getDirections(p)
	if dirs[0] == d {
		return dirs[1]
	}
	return dirs[0]
}
