// Package must contains helper functions that panic if they have errors.
package must

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/glennhartmann/aoc23/src/common"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		common.Panicf("invalid int for Atoi: %s", s)
	}
	return i
}

func Atoi64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		common.Panicf("invalid int64 for Atoi64: %s", s)
	}
	return i
}

func ForEachLineOfStreamedInput(f func(lineNum int, s string)) {
	r := bufio.NewReader(os.Stdin)
	lineNum := 0
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			log.Printf("EOF")
			break
		}
		if err != nil {
			common.Panicf("unable to read from stdin: %v", err)
		}
		s = strings.TrimSuffix(s, "\n")
		log.Printf("current line: %q", s)

		f(lineNum, s)

		lineNum++
	}
}

func GetFullInput() []string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		common.Panicf("error reading from stdin: %v", err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}
