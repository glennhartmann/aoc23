package common_10

import (
	"log"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/grid/d4"

	c22 "github.com/glennhartmann/aoc22/src/common"
)

// part 1 main
func GetFarthestDistance(lines []string) int {
	lines = common.AddSentinal(lines, ".")

	y, x := d4.MustFindInStringGrid(lines, 'S')
	s := DetermineStartShape(lines, x, y)
	log.Printf("start [%d, %d] is %c", x, y, s)

	nodesVisited := 1
	dirs := GetDirections(s)
	dir := dirs[0]
	for {
		y, x = d4.GetNextCell(y, x, dir)
		log.Printf("moving %v to [%d, %d]", dir, x, y)

		if lines[y][x] == 'S' {
			break
		}
		nodesVisited++
		dir = GetOtherDir(lines[y][x], d4.OppositeDir(dir))
	}
	log.Printf("visited %d nodes", nodesVisited)
	log.Printf("farthest distance: %d", nodesVisited/2)
	return nodesVisited / 2
}

// part 2 main
func GetEnclosedAreas(lines []string) (int, int) {
	m := make([][]byte, len(lines))
	for i := 0; i < len(m); i++ {
		m[i] = []byte(c22.Padding("?", len(lines[0])))
	}

	lines = common.AddSentinal(lines, ".")
	m = common.AddSentinal2(m, '.')

	y, x := d4.MustFindInStringGrid(lines, 'S')
	s := DetermineStartShape(lines, x, y)
	log.Printf("start [%d, %d] is %c", x, y, s)

	m[y][x] = 'S'

	// trace main loop
	dirs := GetDirections(s)
	dir := dirs[0]
	for {
		y, x = d4.GetNextCell(y, x, dir)
		log.Printf("moving %v to [%d, %d]", dir, x, y)

		if lines[y][x] == 'S' {
			break
		}

		m[y][x] = 'M' // part of Main loop

		dir = GetOtherDir(lines[y][x], d4.OppositeDir(dir))
	}

	log.Printf("secondary map main loop:\n%s", Consolify(m))

	// categorize tiles adjacent to main loop as side 1 or side 2
	// (we cheat and don't bother determining which side is 'in'
	// and which is 'out')
	y, x = d4.MustFindInStringGrid(lines, 'S')
	dir = dirs[0]
	side1Dir := GetStartSide1Dir(s, dir)
	adjacent := FindAdjacentTilesInternal(lines, x, y, side1Dir, s)
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

		y, x = d4.GetNextCell(y, x, dir)

		if lines[y][x] == 'S' {
			break
		}

		oldDir := dir
		dir = GetOtherDir(lines[y][x], d4.OppositeDir(dir))
		side1Dir = GetSide1Dir(oldDir, side1Dir, lines[y][x])
		adjacent = FindAdjacentTiles(lines, x, y, side1Dir)
	}

	log.Printf("secondary map adjacent tiles:\n%s", Consolify(m))

	// flood fill remaining ?s
	for row := range m {
		for col := range m[row] {
			resolve(m, row, col)
		}
	}

	log.Printf("secondary map full:\n%s", Consolify(m))

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
	return counts[0], counts[1]
}

func DetermineStartShape(lines []string, x, y int) byte {
	cl, cr, cu, cd := ConnectsRight(lines[y][x-1]), ConnectsLeft(lines[y][x+1]), ConnectsDown(lines[y-1][x]), ConnectsUp(lines[y+1][x])
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

func ConnectsLeft(p byte) bool {
	return p == '-' || p == 'J' || p == '7'
}

func ConnectsRight(p byte) bool {
	return p == '-' || p == 'L' || p == 'F'
}

func ConnectsUp(p byte) bool {
	return p == '|' || p == 'L' || p == 'J'
}

func ConnectsDown(p byte) bool {
	return p == '|' || p == '7' || p == 'F'
}

func GetDirections(p byte) []d4.Direction {
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

func GetOtherDir(p byte, d d4.Direction) d4.Direction {
	dirs := GetDirections(p)
	if dirs[0] == d {
		return dirs[1]
	}
	return dirs[0]
}

func Consolify(m [][]byte) string {
	return c22.Fjoin(m, "\n", func(e []byte) string { return string(e) })
}

type AdjacentTiles struct{ side1, side2 []Point }
type Point struct{ x, y int }

func FindAdjacentTiles(lines []string, x, y int, d d4.Direction) AdjacentTiles {
	return FindAdjacentTilesInternal(lines, x, y, d, lines[y][x])
}

func FindAdjacentTilesInternal(lines []string, x, y int, d d4.Direction, p byte) AdjacentTiles {
	adj := AdjacentTiles{}
	switch d {
	case d4.Down:
		switch p {
		case '|':
			break
		case '-':
			/*
			 * ?2?
			 * ?-?
			 * ?1?
			 */
			adj.side1 = append(adj.side1, Point{x, y + 1})
			adj.side2 = append(adj.side2, Point{x, y - 1})
		case '7':
			/*
			 * ?22
			 * ?72
			 * 1??
			 */
			adj.side1 = append(adj.side1, Point{x - 1, y + 1})
			adj.side2 = append(adj.side2, Point{x, y - 1}, Point{x + 1, y - 1}, Point{x + 1, y})
		case 'F':
			/*
			 * 22?
			 * 2F?
			 * ??1
			 */
			adj.side1 = append(adj.side1, Point{x + 1, y + 1})
			adj.side2 = append(adj.side2, Point{x, y - 1}, Point{x - 1, y - 1}, Point{x - 1, y})
		case 'L':
			/*
			 * ??2
			 * 1L?
			 * 11?
			 */
			adj.side1 = append(adj.side1, Point{x, y + 1}, Point{x - 1, y + 1}, Point{x - 1, y})
			adj.side2 = append(adj.side2, Point{x + 1, y - 1})
		case 'J':
			/*
			 * 2??
			 * ?J1
			 * ?11
			 */
			adj.side1 = append(adj.side1, Point{x, y + 1}, Point{x + 1, y + 1}, Point{x + 1, y})
			adj.side2 = append(adj.side2, Point{x - 1, y - 1})
		default:
			panic("bad pipe in FindAdjacentTiles down")
		}
	case d4.Up:
		switch p {
		case '|':
			break
		case '-':
			/*
			 * ?1?
			 * ?-?
			 * ?2?
			 */
			adj.side1 = append(adj.side1, Point{x, y - 1})
			adj.side2 = append(adj.side2, Point{x, y + 1})
		case 'L':
			/*
			 * ??1
			 * 2L?
			 * 22?
			 */
			adj.side1 = append(adj.side1, Point{x + 1, y - 1})
			adj.side2 = append(adj.side2, Point{x, y + 1}, Point{x - 1, y + 1}, Point{x - 1, y})
		case 'J':
			/*
			 * 1??
			 * ?J2
			 * ?22
			 */
			adj.side1 = append(adj.side1, Point{x - 1, y - 1})
			adj.side2 = append(adj.side2, Point{x, y + 1}, Point{x + 1, y + 1}, Point{x + 1, y})
		case 'F':
			/*
			 * 11?
			 * 1F?
			 * ??2
			 */
			adj.side1 = append(adj.side1, Point{x, y - 1}, Point{x - 1, y - 1}, Point{x - 1, y})
			adj.side2 = append(adj.side2, Point{x + 1, y + 1})
		case '7':
			/*
			 * ?11
			 * ?71
			 * 2??
			 */
			adj.side1 = append(adj.side1, Point{x, y - 1}, Point{x + 1, y - 1}, Point{x + 1, y})
			adj.side2 = append(adj.side2, Point{x - 1, y + 1})
		default:
			panic("bad pipe in FindAdjacentTiles up")
		}
	case d4.Left:
		switch p {
		case '-':
			break
		case '|':
			/*
			 * ???
			 * 1|2
			 * ???
			 */
			adj.side1 = append(adj.side1, Point{x - 1, y})
			adj.side2 = append(adj.side2, Point{x + 1, y})
		case '7':
			/*
			 * ?22
			 * ?72
			 * 1??
			 */
			adj.side1 = append(adj.side1, Point{x - 1, y + 1})
			adj.side2 = append(adj.side2, Point{x, y - 1}, Point{x + 1, y - 1}, Point{x + 1, y})
		case 'J':
			/*
			 * 1??
			 * ?J2
			 * ?22
			 */
			adj.side1 = append(adj.side1, Point{x - 1, y - 1})
			adj.side2 = append(adj.side2, Point{x, y + 1}, Point{x + 1, y + 1}, Point{x + 1, y})
		case 'F':
			/*
			 * 11?
			 * 1F?
			 * ??2
			 */
			adj.side1 = append(adj.side1, Point{x, y - 1}, Point{x - 1, y - 1}, Point{x - 1, y})
			adj.side2 = append(adj.side2, Point{x + 1, y + 1})
		case 'L':
			/*
			 * ??2
			 * 1L?
			 * 11?
			 */
			adj.side1 = append(adj.side1, Point{x - 1, y}, Point{x - 1, y + 1}, Point{x, y + 1})
			adj.side2 = append(adj.side2, Point{x + 1, y - 1})
		default:
			panic("bad pipe in FindAdjacentTiles left")
		}
	case d4.Right:
		switch p {
		case '-':
			break
		case '|':
			/*
			 * ???
			 * 2|1
			 * ???
			 */
			adj.side1 = append(adj.side1, Point{x + 1, y})
			adj.side2 = append(adj.side2, Point{x - 1, y})
		case 'F':
			/*
			 * 22?
			 * 2F?
			 * ??1
			 */
			adj.side1 = append(adj.side1, Point{x + 1, y + 1})
			adj.side2 = append(adj.side2, Point{x, y - 1}, Point{x - 1, y - 1}, Point{x - 1, y})
		case 'L':
			/*
			 * ??1
			 * 2L?
			 * 22?
			 */
			adj.side1 = append(adj.side1, Point{x + 1, y - 1})
			adj.side2 = append(adj.side2, Point{x - 1, y}, Point{x - 1, y + 1}, Point{x, y + 1})
		case '7':
			/*
			 * ?11
			 * ?71
			 * 2??
			 */
			adj.side1 = append(adj.side1, Point{x, y - 1}, Point{x + 1, y - 1}, Point{x + 1, y})
			adj.side2 = append(adj.side2, Point{x - 1, y + 1})
		case 'J':
			/*
			 * 2??
			 * ?J1
			 * ?11
			 */
			adj.side1 = append(adj.side1, Point{x + 1, y}, Point{x + 1, y + 1}, Point{x, y + 1})
			adj.side2 = append(adj.side2, Point{x - 1, y - 1})
		default:
			panic("bad pipe in FindAdjacentTiles right")
		}
	default:
		panic("bad direction in FindAdjacentTiles")
	}

	return adj
}

func GetSide1Dir(oldDir, oldSide1Dir d4.Direction, p byte) d4.Direction {
	switch p {
	case '-', '|':
		return oldSide1Dir
	case 'L':
		if oldDir == d4.Down {
			if oldSide1Dir == d4.Right {
				return d4.Up
			}
			return d4.Down
		}
		if oldSide1Dir == d4.Up {
			return d4.Right
		}
		return d4.Left
	case 'J':
		if oldDir == d4.Down {
			if oldSide1Dir == d4.Right {
				return d4.Down
			}
			return d4.Up
		}
		if oldSide1Dir == d4.Up {
			return d4.Left
		}
		return d4.Right
	case '7':
		if oldDir == d4.Up {
			if oldSide1Dir == d4.Right {
				return d4.Up
			}
			return d4.Down
		}
		if oldSide1Dir == d4.Up {
			return d4.Right
		}
		return d4.Left
	case 'F':
		if oldDir == d4.Up {
			if oldSide1Dir == d4.Right {
				return d4.Down
			}
			return d4.Up
		}
		if oldSide1Dir == d4.Up {
			return d4.Left
		}
		return d4.Right
	default:
		panic("bad pipe in getSideDir")
	}
}

func GetStartSide1Dir(p byte, d d4.Direction) d4.Direction {
	switch p {
	case '|':
		return d4.Left
	case '-':
		return d4.Up
	case 'L':
		if d == d4.Right {
			return d4.Up
		}
		return d4.Right
	case 'J':
		if d == d4.Left {
			return d4.Up
		}
		return d4.Left
	case '7':
		if d == d4.Left {
			return d4.Down
		}
		return d4.Left
	case 'F':
		if d == d4.Right {
			return d4.Down
		}
		return d4.Right
	default:
		panic("bad pipe in GetStartSide1Dir")
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
