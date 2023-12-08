package main

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/must"
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
		asteriskIdx := -1
		for {
			asteriskIdx = getNextAsterisk(line, asteriskIdx+1)
			if asteriskIdx == -1 {
				break
			}

			log.Printf("line %d: found asterisk at index %d", lineIndex, asteriskIdx)

			adj := getAdjacentNumbers(lines, lineIndex, asteriskIdx)
			if len(adj) == 2 {
				ratio := adj[0] * adj[1]
				log.Printf("line %d: asterisk at index %d is a gear with part numbers %d and %d and ratio %d", lineIndex, asteriskIdx, adj[0], adj[1], ratio)
				sum += ratio
			}
		}
		log.Printf("running sum after line %d: %d", lineIndex, sum)
	}

	log.Printf("final sum: %d", sum)
}

func getNextAsterisk(line string, scanStart int) int {
	idx := strings.Index(line[scanStart:], "*")
	if idx == -1 {
		return -1
	}
	return idx + scanStart
}

func getAdjacentNumbers(lines []string, lineIndex, col int) []int {
	adjacencies := []int{}

	if lineIndex > 0 {
		adjacencies = append(adjacencies, topBottomAdjacencies(lines[lineIndex-1], col)...)
	}
	if lineIndex < len(lines)-1 {
		adjacencies = append(adjacencies, topBottomAdjacencies(lines[lineIndex+1], col)...)
	}

	return append(adjacencies, leftRightAdjacencies(lines[lineIndex], col)...)
}

func topBottomAdjacencies(line string, col int) []int {
	if common.IsDigit(line[col]) {
		return []int{numberSurrounding(line, col)}
	}

	return leftRightAdjacencies(line, col)
}

func leftRightAdjacencies(line string, col int) []int {
	nums := []int{}
	if col > 0 && common.IsDigit(line[col-1]) {
		nums = append(nums, numberSurrounding(line, col-1))
	}
	if col < len(line)-1 && common.IsDigit(line[col+1]) {
		nums = append(nums, numberSurrounding(line, col+1))
	}

	return nums
}

func numberSurrounding(line string, col int) int {
	left, right := col, col

	for i := col; i >= 0; i-- {
		if common.IsDigit(line[i]) {
			left = i
		} else {
			break
		}
	}

	for i := col; i < len(line); i++ {
		if common.IsDigit(line[i]) {
			right = i
		} else {
			break
		}
	}

	return must.Atoi(line[left : right+1])
}
