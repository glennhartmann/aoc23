package main

import (
	"log"

	"github.com/glennhartmann/aoclib/common"
	"github.com/glennhartmann/aoclib/must"
)

func main() {
	lines := must.GetFullInput()

	sum := 0
	for lineIndex, line := range lines {
		start, end := -1, -1
		for {
			start, end = getNextNumber(line, end+1)
			if start == -1 {
				break
			}

			numStr := line[start : end+1]
			num := must.Atoi(numStr)
			log.Printf("line %d: found number: %d", lineIndex, num)

			if isAdjacentToSymbol(lines, lineIndex, start, end) {
				log.Printf("line %d: %d is adjacent to a symbol", lineIndex, num)
				sum += num
			}
		}
		log.Printf("running sum after line %d: %d", lineIndex, sum)
	}

	log.Printf("final sum: %d", sum)
}

func getNextNumber(line string, scanStart int) (numberStart, numberEnd int) {
	if scanStart >= len(line) {
		return -1, -1
	}
	numberStart, numberEnd = -1, -1

	inNumber := false
	for i := scanStart; i < len(line); i++ {
		if inNumber {
			if common.IsDigit(line[i]) {
				numberEnd = i
			} else {
				break
			}
		} else if common.IsDigit(line[i]) {
			numberStart = i
			numberEnd = i
			inNumber = true
		}
	}

	return numberStart, numberEnd
}

func isAdjacentToSymbol(lines []string, lineIndex, start, end int) bool {
	top := max(0, lineIndex-1)
	left := max(0, start-1)
	bottom := min(lineIndex+1, len(lines)-1)
	right := min(end+1, len(lines[lineIndex])-1)

	for row := top; row <= bottom; row++ {
		for col := left; col <= right; col++ {
			if isSymbol(lines[row][col]) {
				return true
			}
		}
	}

	return false
}

func isSymbol(b byte) bool {
	return b != '.' && !common.IsDigit(b)
}
