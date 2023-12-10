package main

import (
	"log"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/must"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	default:
		panic("bad direction for stringer")
	}
}

func main() {
	lines := must.GetFullInput()
	lines = common.AddSentinal(lines, ".")

	x, y := findStart(lines)
	s := determineStartShape(lines, x, y)
	log.Printf("start [%d, %d] is %c", x, y, s)

	nodesVisited := 1
	dirs := getDirections(s)
	dir := dirs[0]
	for {
		x, y = move(x, y, dir)
		log.Printf("moving %v to [%d, %d]", dir, x, y)

		if lines[y][x] == 'S' {
			break
		}
		nodesVisited++
		dir = getOtherDir(lines[y][x], oppositeDir(dir))
	}
	log.Printf("visited %d nodes", nodesVisited)
	log.Printf("farthest distance: %d", nodesVisited/2)
}

func findStart(lines []string) (x, y int) {
	for row := range lines {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] == 'S' {
				return col, row
			}
		}
	}
	panic("start not found")
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

func getDirections(p byte) []direction {
	switch p {
	case '|':
		return []direction{up, down}
	case '-':
		return []direction{left, right}
	case 'L':
		return []direction{up, right}
	case 'J':
		return []direction{up, left}
	case '7':
		return []direction{down, left}
	case 'F':
		return []direction{down, right}
	default:
		panic("bad direction")
	}
}

func move(x, y int, d direction) (xNew, yNew int) {
	switch d {
	case up:
		return x, y - 1
	case down:
		return x, y + 1
	case left:
		return x - 1, y
	case right:
		return x + 1, y
	default:
		panic("bad direction for move")
	}
}

func oppositeDir(dir direction) direction {
	switch dir {
	case up:
		return down
	case down:
		return up
	case left:
		return right
	case right:
		return left
	default:
		panic("bad direction for opposite")
	}
}

func getOtherDir(p byte, d direction) direction {
	dirs := getDirections(p)
	if dirs[0] == d {
		return dirs[1]
	}
	return dirs[0]
}
