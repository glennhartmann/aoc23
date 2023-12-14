package main

import (
	"log"

	"github.com/glennhartmann/aoc23/src/common/must"
)

func main() {
	lines := must.GetFullInputAsBytes()

	for c := 0; c < len(lines[0]); c++ {
		slot := 0
		rounds := 0
		for r := 0; r < len(lines); r++ {
			if lines[r][c] == '#' {
				moveUp(lines, c, slot, rounds)
				rounds = 0
				slot = r + 1
			} else if lines[r][c] == 'O' {
				rounds++
				lines[r][c] = '.'
			}
		}
		moveUp(lines, c, slot, rounds)
	}

	log.Printf("new layout:")
	for r := range lines {
		log.Printf("%s", string(lines[r]))
	}

	sum := 0
	for r := range lines {
		sum += countRow(lines[r]) * (len(lines) - r)
	}
	log.Printf("sum: %d", sum)
}

func moveUp(lines [][]byte, c, slot, rounds int) {
	for i := 0; i < rounds; i++ {
		lines[slot+i][c] = 'O'
	}
}

func countRow(line []byte) int {
	count := 0
	for c := range line {
		if line[c] == 'O' {
			count++
		}
	}
	return count
}
