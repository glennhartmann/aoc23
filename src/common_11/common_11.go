package common_11

import (
	"log"

	"github.com/glennhartmann/aoc23/src/common/must"

	c22 "github.com/glennhartmann/aoc22/src/common"
)

type point struct {
	r, c int
}

func Compute(expansionFactor int) {
	lines := must.GetFullInput()

	rowIsEmpty := make([]bool, len(lines[0]))
	for i := range rowIsEmpty {
		rowIsEmpty[i] = true
	}

	colIsEmpty := make([]bool, len(lines))
	for i := range colIsEmpty {
		colIsEmpty[i] = true
	}

	for row := range lines {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] == '#' {
				rowIsEmpty[row] = false
				colIsEmpty[col] = false
			}
		}
	}

	galaxies := make([]point, 0, 50)

	virtualRow := 0
	virtualCol := 0
	for row := range lines {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] == '#' {
				galaxies = append(galaxies, point{r: virtualRow, c: virtualCol})
			}
			if colIsEmpty[col] {
				virtualCol += expansionFactor - 1
			}
			virtualCol++
		}
		if rowIsEmpty[row] {
			virtualRow += expansionFactor - 1
		}
		virtualCol = 0
		virtualRow++
	}

	log.Printf("galaxy coordinates after expansion: %+v", galaxies)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			d := dist(galaxies[i], galaxies[j])
			log.Printf("distance between galaxy %d and galaxy %d: %d", i, j, d)
			sum += d
		}
	}
	log.Printf("sum: %d", sum)
}

func dist(a, b point) int {
	return c22.Abs(a.r-b.r) + c22.Abs(a.c-b.c)
}
