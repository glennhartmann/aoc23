package common

import (
	"fmt"

	c22 "github.com/glennhartmann/aoc22/src/common"
)

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func DigitToInt(b byte) int {
	return int(b - '0')
}

func Panicf(fmtStr string, args ...any) {
	panic(fmt.Sprintf(fmtStr, args...))
}

func AddSentinal(lines []string, c string) []string {
	line := c22.Padding(c, len(lines[0]))
	lines = append([]string{line}, lines...)
	lines = append(lines, line)

	for i := range lines {
		lines[i] = fmt.Sprintf("%s%s%s", c, lines[i], c)
	}

	return lines
}

func AddSentinal2[T any](lines [][]T, c T) [][]T {
	line := make([]T, len(lines[0]))
	for i := range line {
		line[i] = c
	}
	lines = append([][]T{line}, lines...)
	lines = append(lines, line)

	for i := range lines {
		lines[i] = append([]T{c}, lines[i]...)
		lines[i] = append(lines[i], c)
	}

	return lines
}
