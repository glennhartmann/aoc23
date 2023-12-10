package main

import (
	"log"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/must"

	c22 "github.com/glennhartmann/aoc22/src/common"
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

	m := make([][]byte, len(lines))
	for i := 0; i < len(m); i++ {
		m[i] = []byte(c22.Padding("?", len(lines[0])))
	}

	lines = common.AddSentinal(lines, ".")
	m = common.AddSentinal2(m, '.')

	x, y := findStart(lines)
	s := determineStartShape(lines, x, y)
	log.Printf("start [%d, %d] is %c", x, y, s)

	m[y][x] = 'S'

	// trace main loop
	dirs := getDirections(s)
	dir := dirs[0]
	for {
		x, y = move(x, y, dir)
		log.Printf("moving %v to [%d, %d]", dir, x, y)

		if lines[y][x] == 'S' {
			break
		}

		m[y][x] = 'M' // part of Main loop

		dir = getOtherDir(lines[y][x], oppositeDir(dir))
	}

	log.Printf("secondary map main loop:\n%s", consolify(m))

	// categorize tiles adjacent to main loop as side 1 or side 2
	// (we cheat and don't bother determining which side is 'in'
	// and which is 'out')
	x, y = findStart(lines)
	dir = dirs[0]
	side1Dir := getStartSide1Dir(s, dir)
	adjacent := findAdjacentTilesInternal(lines, x, y, side1Dir, s)
	for {
		for _, p := range adjacent.side1 {
			if m[p.y][p.x] == '?' {
				m[p.y][p.x] = '1' // part of side 1
			}
		}
		for _, p := range adjacent.side2 {
			if m[p.y][p.x] == '?' {
				m[p.y][p.x] = '2' // part of side 2
			}
		}

		x, y = move(x, y, dir)

		if lines[y][x] == 'S' {
			break
		}

		oldDir := dir
		dir = getOtherDir(lines[y][x], oppositeDir(dir))
		side1Dir = getSide1Dir(oldDir, side1Dir, lines[y][x])
		adjacent = findAdjacentTiles(lines, x, y, side1Dir)
	}

	log.Printf("secondary map adjacent tiles:\n%s", consolify(m))

	// flood fill remaining ?s
	for row := range m {
		for col := range m[row] {
			resolve(m, row, col)
		}
	}

	log.Printf("secondary map full:\n%s", consolify(m))

	// count both sides
	counts := []int{0, 0}
	ind := func(c byte) int {
		if c == '1' {
			return 0
		}
		return 1
	}
	for row := range m {
		for col := range m[row] {
			if m[row][col] == '1' || m[row][col] == '2' {
				counts[ind(m[row][col])]++
			}
		}
	}
	log.Printf("final counts: side 1 (%d), side 2 (%d)", counts[0], counts[1])
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

func consolify(m [][]byte) string {
	return c22.Fjoin(m, "\n", func(e []byte) string { return string(e) })
}

type adjacentTiles struct{ side1, side2 []point }
type point struct{ x, y int }

func findAdjacentTiles(lines []string, x, y int, d direction) adjacentTiles {
	return findAdjacentTilesInternal(lines, x, y, d, lines[y][x])
}

func findAdjacentTilesInternal(lines []string, x, y int, d direction, p byte) adjacentTiles {
	adj := adjacentTiles{}
	switch d {
	case down:
		switch p {
		case '|':
			break
		case '-':
			/*
			 * ?2?
			 * ?-?
			 * ?1?
			 */
			adj.side1 = append(adj.side1, point{x, y + 1})
			adj.side2 = append(adj.side2, point{x, y - 1})
		case '7':
			/*
			 * ?22
			 * ?72
			 * 1??
			 */
			adj.side1 = append(adj.side1, point{x - 1, y + 1})
			adj.side2 = append(adj.side2, point{x, y - 1}, point{x + 1, y - 1}, point{x + 1, y})
		case 'F':
			/*
			 * 22?
			 * 2F?
			 * ??1
			 */
			adj.side1 = append(adj.side1, point{x + 1, y + 1})
			adj.side2 = append(adj.side2, point{x, y - 1}, point{x - 1, y - 1}, point{x - 1, y})
		case 'L':
			/*
			 * ??2
			 * 1L?
			 * 11?
			 */
			adj.side1 = append(adj.side1, point{x, y + 1}, point{x - 1, y + 1}, point{x - 1, y})
			adj.side2 = append(adj.side2, point{x + 1, y - 1})
		case 'J':
			/*
			 * 2??
			 * ?J1
			 * ?11
			 */
			adj.side1 = append(adj.side1, point{x, y + 1}, point{x + 1, y + 1}, point{x + 1, y})
			adj.side2 = append(adj.side2, point{x - 1, y - 1})
		default:
			panic("bad pipe in findAdjacentTiles down")
		}
	case up:
		switch p {
		case '|':
			break
		case '-':
			/*
			 * ?1?
			 * ?-?
			 * ?2?
			 */
			adj.side1 = append(adj.side1, point{x, y - 1})
			adj.side2 = append(adj.side2, point{x, y + 1})
		case 'L':
			/*
			 * ??1
			 * 2L?
			 * 22?
			 */
			adj.side1 = append(adj.side1, point{x + 1, y - 1})
			adj.side2 = append(adj.side2, point{x, y + 1}, point{x - 1, y + 1}, point{x - 1, y})
		case 'J':
			/*
			 * 1??
			 * ?J2
			 * ?22
			 */
			adj.side1 = append(adj.side1, point{x - 1, y - 1})
			adj.side2 = append(adj.side2, point{x, y + 1}, point{x + 1, y + 1}, point{x + 1, y})
		case 'F':
			/*
			 * 11?
			 * 1F?
			 * ??2
			 */
			adj.side1 = append(adj.side1, point{x, y - 1}, point{x - 1, y - 1}, point{x - 1, y})
			adj.side2 = append(adj.side2, point{x + 1, y + 1})
		case '7':
			/*
			 * ?11
			 * ?71
			 * 2??
			 */
			adj.side1 = append(adj.side1, point{x, y - 1}, point{x + 1, y - 1}, point{x + 1, y})
			adj.side2 = append(adj.side2, point{x - 1, y + 1})
		default:
			panic("bad pipe in findAdjacentTiles up")
		}
	case left:
		switch p {
		case '-':
			break
		case '|':
			/*
			 * ???
			 * 1|2
			 * ???
			 */
			adj.side1 = append(adj.side1, point{x - 1, y})
			adj.side2 = append(adj.side2, point{x + 1, y})
		case '7':
			/*
			 * ?22
			 * ?72
			 * 1??
			 */
			adj.side1 = append(adj.side1, point{x - 1, y + 1})
			adj.side2 = append(adj.side2, point{x, y - 1}, point{x + 1, y - 1}, point{x + 1, y})
		case 'J':
			/*
			 * 1??
			 * ?J2
			 * ?22
			 */
			adj.side1 = append(adj.side1, point{x - 1, y - 1})
			adj.side2 = append(adj.side2, point{x, y + 1}, point{x + 1, y + 1}, point{x + 1, y})
		case 'F':
			/*
			 * 11?
			 * 1F?
			 * ??2
			 */
			adj.side1 = append(adj.side1, point{x, y - 1}, point{x - 1, y - 1}, point{x - 1, y})
			adj.side2 = append(adj.side2, point{x + 1, y + 1})
		case 'L':
			/*
			 * ??2
			 * 1L?
			 * 11?
			 */
			adj.side1 = append(adj.side1, point{x - 1, y}, point{x - 1, y + 1}, point{x, y + 1})
			adj.side2 = append(adj.side2, point{x + 1, y - 1})
		default:
			panic("bad pipe in findAdjacentTiles left")
		}
	case right:
		switch p {
		case '-':
			break
		case '|':
			/*
			 * ???
			 * 2|1
			 * ???
			 */
			adj.side1 = append(adj.side1, point{x + 1, y})
			adj.side2 = append(adj.side2, point{x - 1, y})
		case 'F':
			/*
			 * 22?
			 * 2F?
			 * ??1
			 */
			adj.side1 = append(adj.side1, point{x + 1, y + 1})
			adj.side2 = append(adj.side2, point{x, y - 1}, point{x - 1, y - 1}, point{x - 1, y})
		case 'L':
			/*
			 * ??1
			 * 2L?
			 * 22?
			 */
			adj.side1 = append(adj.side1, point{x + 1, y - 1})
			adj.side2 = append(adj.side2, point{x - 1, y}, point{x - 1, y + 1}, point{x, y + 1})
		case '7':
			/*
			 * ?11
			 * ?71
			 * 2??
			 */
			adj.side1 = append(adj.side1, point{x, y - 1}, point{x + 1, y - 1}, point{x + 1, y})
			adj.side2 = append(adj.side2, point{x - 1, y + 1})
		case 'J':
			/*
			 * 2??
			 * ?J1
			 * ?11
			 */
			adj.side1 = append(adj.side1, point{x + 1, y}, point{x + 1, y + 1}, point{x, y + 1})
			adj.side2 = append(adj.side2, point{x - 1, y - 1})
		default:
			panic("bad pipe in findAdjacentTiles right")
		}
	default:
		panic("bad direction in findAdjacentTiles")
	}

	return adj
}

func getSide1Dir(oldDir, oldSide1Dir direction, p byte) direction {
	switch p {
	case '-', '|':
		return oldSide1Dir
	case 'L':
		if oldDir == down {
			if oldSide1Dir == right {
				return up
			}
			return down
		}
		if oldSide1Dir == up {
			return right
		}
		return left
	case 'J':
		if oldDir == down {
			if oldSide1Dir == right {
				return down
			}
			return up
		}
		if oldSide1Dir == up {
			return left
		}
		return right
	case '7':
		if oldDir == up {
			if oldSide1Dir == right {
				return up
			}
			return down
		}
		if oldSide1Dir == up {
			return right
		}
		return left
	case 'F':
		if oldDir == up {
			if oldSide1Dir == right {
				return down
			}
			return up
		}
		if oldSide1Dir == up {
			return left
		}
		return right
	default:
		panic("bad pipe in getSideDir")
	}
}

func getStartSide1Dir(p byte, d direction) direction {
	switch p {
	case '|':
		return left
	case '-':
		return up
	case 'L':
		if d == right {
			return up
		}
		return right
	case 'J':
		if d == left {
			return up
		}
		return left
	case '7':
		if d == left {
			return down
		}
		return left
	case 'F':
		if d == right {
			return down
		}
		return right
	default:
		panic("bad pipe in getStartSide1Dir")
	}
}

func resolve(m [][]byte, row, col int) {
	visited := make([][]bool, len(m))
	for i := range visited {
		visited[i] = make([]bool, len(m[0]))
	}
	resolveInternal(m, row, col, visited)
}

func resolveInternal(m [][]byte, row, col int, visited [][]bool) {
	if m[row][col] != '?' || visited[row][col] {
		return
	}
	visited[row][col] = true

	resolveInternal(m, row+1, col, visited)
	if m[row+1][col] == '1' || m[row+1][col] == '2' {
		m[row][col] = m[row+1][col]
		return
	}

	resolveInternal(m, row, col+1, visited)
	if m[row][col+1] == '1' || m[row][col+1] == '2' {
		m[row][col] = m[row][col+1]
		return
	}

	resolveInternal(m, row-1, col, visited)
	if m[row-1][col] == '1' || m[row-1][col] == '2' {
		m[row][col] = m[row-1][col]
		return
	}

	resolveInternal(m, row, col-1, visited)
	if m[row][col-1] == '1' || m[row][col-1] == '2' {
		m[row][col] = m[row][col-1]
		return
	}
}
