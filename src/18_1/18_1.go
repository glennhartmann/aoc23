package main

import (
	"log"
	"regexp"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/grid/d4"
	"github.com/glennhartmann/aoc23/src/common/must"
	c10 "github.com/glennhartmann/aoc23/src/common_10"
)

var lineRx = regexp.MustCompile(`(U|D|L|R) (\d+) \((#[[:xdigit:]]{6})\)`)

type dig struct {
	dir    d4.Direction
	length int
	colour string
}

func main() {
	lines := must.GetFullInput()

	grid := make([][]byte, 400)
	for i := range grid {
		grid[i] = make([]byte, 800)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	digs := make([]dig, 0, len(lines))
	for _, line := range lines {
		match := must.FindStringSubmatch(lineRx, line, 4)
		digs = append(digs, dig{
			dir:    d4.DirForUDLR(match[1]),
			length: must.Atoi(match[2]),
			colour: match[3],
		})
	}

	previousDir := digs[len(digs)-1].dir
	r, c := 250, 250
	borderLength := 0
	for i, dig := range digs {
		log.Printf("%d: %+v", i, dig)
		borderLength += dig.length

		for ln := 0; ln < dig.length; ln++ {
			grid[r][c] = getSymbol(previousDir, dig.dir)
			previousDir = dig.dir
			r, c = d4.GetNextCell(r, c, dig.dir)
		}
	}
	grid[r][c] = 'S'

	enclosedSize1, enclosedSize2 := c10.GetEnclosedAreas(common.ByteSlice2ToStringSlice(grid))
	log.Printf("final counts: %d, %d", borderLength+enclosedSize1, borderLength+enclosedSize2)
}

func getSymbol(prevDir, dir d4.Direction) byte {
	switch {
	case (prevDir == d4.Up || prevDir == d4.Down) && (dir == d4.Up || dir == d4.Down):
		return '|'
	case (prevDir == d4.Left || prevDir == d4.Right) && (dir == d4.Left || dir == d4.Right):
		return '-'
	case (prevDir == d4.Down && dir == d4.Right) || (prevDir == d4.Left && dir == d4.Up):
		return 'L'
	case (prevDir == d4.Down && dir == d4.Left) || (prevDir == d4.Right && dir == d4.Up):
		return 'J'
	case (prevDir == d4.Right && dir == d4.Down) || (prevDir == d4.Up && dir == d4.Left):
		return '7'
	case (prevDir == d4.Left && dir == d4.Down) || (prevDir == d4.Up && dir == d4.Right):
		return 'F'
	default:
		panic("invalid combination of previous and current directions")
	}
}
