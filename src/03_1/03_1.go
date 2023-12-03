package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("error reading from stdin")
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sum := 0
	for lineIndex, line := range lines {
		start, end := -1, -1
		for {
			start, end = getNextNumber(line, end+1)
			if start == -1 {
				break
			}

			numStr := line[start : end+1]
			num := val(numStr)
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
			if isDigit(line[i]) {
				numberEnd = i
			} else {
				break
			}
		} else if isDigit(line[i]) {
			numberStart = i
			numberEnd = i
			inNumber = true
		}
	}

	return numberStart, numberEnd
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isSymbol(b byte) bool {
	return b != '.' && !isDigit(b)
}

func val(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("bad strconv")
	}
	return i
}
