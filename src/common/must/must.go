// Package must contains helper functions that panic if they have errors.
package must

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
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

func FindStringSubmatch(rx *regexp.Regexp, s string, expectedLen int) []string {
	m := rx.FindStringSubmatch(s)
	if len(m) != expectedLen {
		common.Panicf("regexp match returned len %d, wanted %d", len(m), expectedLen)
	}
	return m
}

func parseListOfNumbersBase[T any](s string, atoi func(string) T) []T {
	sp := strings.Split(s, " ")
	ret := make([]T, 0, len(sp))
	for _, i := range sp {
		if i == "" {
			continue
		}
		ret = append(ret, atoi(i))
	}
	return ret
}

func ParseListOfNumbers(s string) []int {
	return parseListOfNumbersBase(s, Atoi)
}

func ParseListOfNumbers64(s string) []int64 {
	return parseListOfNumbersBase(s, Atoi64)
}
