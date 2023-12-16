package main

import (
	"fmt"
	"log"
	"strings"

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

func (dir direction) String() string {
	switch dir {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	default:
		panic("bad direction")
	}
}

type beam struct {
	dir  direction
	r, c int
}

type cell struct {
	beams map[direction]struct{}
}

func main() {
	lines := must.GetFullInput()
	lines = common.AddSentinal(lines, ",")

	beams := make([]*beam, 0, 10)
	beams = append(beams, &beam{dir: right, r: 1, c: 1})

	cells := make([][]cell, len(lines))
	for r := range cells {
		cells[r] = make([]cell, len(lines[0]))
		for c := range cells[r] {
			cells[r][c] = cell{beams: make(map[direction]struct{}, 4)}
		}
	}
	for b := 0; b < len(beams); b++ {
		for {
			r, c, dir := beams[b].r, beams[b].c, beams[b].dir
			if _, ok := cells[r][c].beams[dir]; ok {
				log.Printf("already saw dir %v on cell {%d, %d} - ending beam %d", dir, r, c, b)
				break
			}
			cells[r][c].beams[dir] = struct{}{}

			if lines[r][c] == ',' {
				log.Printf("on sentinal {%d, %d} - ending beam %d", r, c, b)
				break
			}

			nDir, nBeam := getNextDir(r, c, lines, dir)
			beams[b].dir = nDir
			log.Printf("beam %d: next dir: %v", b, nDir)
			beams[b].r, beams[b].c = getNextCell(r, c, nDir)
			log.Printf("beam %d: next cell: {%d, %d}", b, beams[b].r, beams[b].c)

			if nBeam != nil {
				nBeam.r, nBeam.c = getNextCell(nBeam.r, nBeam.c, nBeam.dir)
				log.Printf("new beam! next cell: {%d, %d}", nBeam.r, nBeam.c)
				beams = append(beams, nBeam)
			}
		}
	}

	log.Printf("beam state:")
	for r := range cells {
		var sb strings.Builder
		for c := range cells[r] {
			if lines[r][c] == '/' || lines[r][c] == '\\' || lines[r][c] == '|' || lines[r][c] == '-' {
				fmt.Fprintf(&sb, "%c", lines[r][c])
			} else if len(cells[r][c].beams) == 1 {
				if lines[r][c] == ',' {
					fmt.Fprintf(&sb, ",")
				} else if _, ok := cells[r][c].beams[up]; ok {
					fmt.Fprintf(&sb, "^")
				} else if _, ok := cells[r][c].beams[down]; ok {
					fmt.Fprintf(&sb, "v")
				} else if _, ok := cells[r][c].beams[left]; ok {
					fmt.Fprintf(&sb, "<")
				} else if _, ok := cells[r][c].beams[right]; ok {
					fmt.Fprintf(&sb, ">")
				}
			} else if len(cells[r][c].beams) > 1 {
				fmt.Fprintf(&sb, "%d", len(cells[r][c].beams))
			} else {
				fmt.Fprintf(&sb, "%c", lines[r][c])
			}
		}
		log.Printf("%s", sb.String())
	}

	log.Printf("energized state:")
	count := 0
	for r := range cells {
		var sb strings.Builder
		for c := range cells[r] {
			if lines[r][c] == ',' {
				fmt.Fprintf(&sb, ",")
			} else if len(cells[r][c].beams) == 0 {
				fmt.Fprintf(&sb, ".")
			} else {
				fmt.Fprintf(&sb, "#")
				count++
			}
		}
		log.Printf("%s", sb.String())
	}
	log.Printf("total energized: %d", count)
}

func getNextDir(r, c int, lines []string, dir direction) (direction, *beam) {
	switch lines[r][c] {
	case '.':
		return dir, nil
	case '/':
		switch dir {
		case up:
			return right, nil
		case down:
			return left, nil
		case left:
			return down, nil
		case right:
			return up, nil
		default:
			panic("bad dir")
		}
	case '\\':
		switch dir {
		case up:
			return left, nil
		case down:
			return right, nil
		case left:
			return up, nil
		case right:
			return down, nil
		default:
			panic("bad dir")
		}
	case '|':
		switch dir {
		case up, down:
			return dir, nil
		case left, right:
			return down, &beam{dir: up, r: r, c: c}
		default:
			panic("bad dir")
		}
	case '-':
		switch dir {
		case up, down:
			return left, &beam{dir: right, r: r, c: c}
		case left, right:
			return dir, nil
		default:
			panic("bad dir")
		}
	default:
		panic("bad cell")
	}
}

func getNextCell(r, c int, dir direction) (nx, ny int) {
	switch dir {
	case up:
		return r - 1, c
	case down:
		return r + 1, c
	case left:
		return r, c - 1
	case right:
		return r, c + 1
	default:
		panic("bad dir")
	}
}
