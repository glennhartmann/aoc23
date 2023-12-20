package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/grid/d4"
	"github.com/glennhartmann/aoc23/src/common/must"
)

type beam struct {
	dir  d4.Direction
	r, c int
}

type cell struct {
	beams map[d4.Direction]struct{}
}

func main() {
	lines := must.GetFullInput()
	lines = common.AddSentinal(lines, ",")

	beams := make([]*beam, 0, 10)
	beams = append(beams, &beam{dir: d4.Right, r: 1, c: 1})

	cells := make([][]cell, len(lines))
	for r := range cells {
		cells[r] = make([]cell, len(lines[0]))
		for c := range cells[r] {
			cells[r][c] = cell{beams: make(map[d4.Direction]struct{}, 4)}
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
			beams[b].r, beams[b].c = d4.GetNextCell(r, c, nDir)
			log.Printf("beam %d: next cell: {%d, %d}", b, beams[b].r, beams[b].c)

			if nBeam != nil {
				nBeam.r, nBeam.c = d4.GetNextCell(nBeam.r, nBeam.c, nBeam.dir)
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
				} else if _, ok := cells[r][c].beams[d4.Up]; ok {
					fmt.Fprintf(&sb, "^")
				} else if _, ok := cells[r][c].beams[d4.Down]; ok {
					fmt.Fprintf(&sb, "v")
				} else if _, ok := cells[r][c].beams[d4.Left]; ok {
					fmt.Fprintf(&sb, "<")
				} else if _, ok := cells[r][c].beams[d4.Right]; ok {
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

func getNextDir(r, c int, lines []string, dir d4.Direction) (d4.Direction, *beam) {
	switch lines[r][c] {
	case '.':
		return dir, nil
	case '/':
		switch dir {
		case d4.Up:
			return d4.Right, nil
		case d4.Down:
			return d4.Left, nil
		case d4.Left:
			return d4.Down, nil
		case d4.Right:
			return d4.Up, nil
		default:
			panic("bad dir")
		}
	case '\\':
		switch dir {
		case d4.Up:
			return d4.Left, nil
		case d4.Down:
			return d4.Right, nil
		case d4.Left:
			return d4.Up, nil
		case d4.Right:
			return d4.Down, nil
		default:
			panic("bad dir")
		}
	case '|':
		switch dir {
		case d4.Up, d4.Down:
			return dir, nil
		case d4.Left, d4.Right:
			return d4.Down, &beam{dir: d4.Up, r: r, c: c}
		default:
			panic("bad dir")
		}
	case '-':
		switch dir {
		case d4.Up, d4.Down:
			return d4.Left, &beam{dir: d4.Right, r: r, c: c}
		case d4.Left, d4.Right:
			return dir, nil
		default:
			panic("bad dir")
		}
	default:
		panic("bad cell")
	}
}
