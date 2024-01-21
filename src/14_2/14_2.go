package main

import (
	"crypto/sha512"
	"log"

	"github.com/glennhartmann/aoclib/must"
)

func main() {
	lines := must.GetFullInputAsBytes()

	m := make(map[string]int, 100)
	cycle := 0
	prev := 0
	for ; ; cycle++ {
		tiltUp(lines)
		tiltLeft(lines)
		tiltDown(lines)
		tiltRight(lines)

		h := hash(lines)
		if _, ok := m[h]; ok {
			prev = m[h]
			break
		}
		m[h] = cycle
	}

	log.Printf("found repeated state at cycle %d (previously saw in cycle %d)", cycle, prev)
	period := cycle - prev
	log.Printf("period = %d", period)

	cycles := 1000000000 - cycle - 1
	desiredCycle := cycles % period
	log.Printf("want pattern in %d cycles", desiredCycle)

	for c := 0; c < desiredCycle; c++ {
		tiltUp(lines)
		tiltLeft(lines)
		tiltDown(lines)
		tiltRight(lines)
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

func tiltUp(lines [][]byte) {
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
}

// yes, I really am just copy/pasting. I don't feel like doing better right now.
func tiltLeft(lines [][]byte) {
	for r := 0; r < len(lines); r++ {
		slot := 0
		rounds := 0
		for c := 0; c < len(lines[0]); c++ {
			if lines[r][c] == '#' {
				moveLeft(lines, r, slot, rounds)
				rounds = 0
				slot = c + 1
			} else if lines[r][c] == 'O' {
				rounds++
				lines[r][c] = '.'
			}
		}
		moveLeft(lines, r, slot, rounds)
	}
}

func tiltDown(lines [][]byte) {
	for c := 0; c < len(lines[0]); c++ {
		slot := len(lines) - 1
		rounds := 0
		for r := len(lines) - 1; r >= 0; r-- {
			if lines[r][c] == '#' {
				moveDown(lines, c, slot, rounds)
				rounds = 0
				slot = r - 1
			} else if lines[r][c] == 'O' {
				rounds++
				lines[r][c] = '.'
			}
		}
		moveDown(lines, c, slot, rounds)
	}
}

func tiltRight(lines [][]byte) {
	for r := 0; r < len(lines); r++ {
		slot := len(lines[0]) - 1
		rounds := 0
		for c := len(lines[0]) - 1; c >= 0; c-- {
			if lines[r][c] == '#' {
				moveRight(lines, r, slot, rounds)
				rounds = 0
				slot = c - 1
			} else if lines[r][c] == 'O' {
				rounds++
				lines[r][c] = '.'
			}
		}
		moveRight(lines, r, slot, rounds)
	}
}

func moveUp(lines [][]byte, c, slot, rounds int) {
	for i := 0; i < rounds; i++ {
		lines[slot+i][c] = 'O'
	}
}

func moveLeft(lines [][]byte, r, slot, rounds int) {
	for i := 0; i < rounds; i++ {
		lines[r][slot+i] = 'O'
	}
}

func moveDown(lines [][]byte, c, slot, rounds int) {
	for i := 0; i < rounds; i++ {
		lines[slot-i][c] = 'O'
	}
}

func moveRight(lines [][]byte, r, slot, rounds int) {
	for i := 0; i < rounds; i++ {
		lines[r][slot-i] = 'O'
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

func hash(lines [][]byte) string {
	b := make([]byte, 0, len(lines)*len(lines[0]))
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			b = append(b, lines[r][c])
		}
	}
	return stringHash(sha512.Sum512(b))
}

func stringHash(b [64]byte) string {
	c := make([]byte, 64)
	for i := 0; i < 64; i++ {
		c[i] = b[i]
	}
	return string(c)
}
