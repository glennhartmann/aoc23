// Package must contains helper functions that panic if they have errors.
package must

import (
	"bufio"
	"encoding/json"
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
		common.Panicf("invalid int for Atoi: %s (%v)", s, err)
	}
	return i
}

func Atoi64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		common.Panicf("invalid int64 for Atoi64: %s (%v)", s, err)
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

func GetFullInputAsBytes() [][]byte {
	strs := GetFullInput()
	ret := make([][]byte, len(strs))
	for i := range ret {
		ret[i] = []byte(strs[i])
	}
	return ret
}

func FindStringSubmatch(rx *regexp.Regexp, s string, expectedLen int) []string {
	m := rx.FindStringSubmatch(s)
	if len(m) != expectedLen {
		common.Panicf("regexp match returned len %d, wanted %d", len(m), expectedLen)
	}
	return m
}

func parseListOfNumbersBase[T any](s, sep string, atoi func(string) T) []T {
	sp := strings.Split(s, sep)
	ret := make([]T, 0, len(sp))
	for _, i := range sp {
		if i == "" {
			continue
		}
		ret = append(ret, atoi(i))
	}
	return ret
}

func ParseListOfNumbers(s, sep string) []int {
	return parseListOfNumbersBase(s, sep, Atoi)
}

func ParseListOfNumbers64(s, sep string) []int64 {
	return parseListOfNumbersBase(s, sep, Atoi64)
}

func JSONMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		common.Panicf("json marshal failed: %v", err)
	}
	return b
}

func JSONMarshalIndent(v any, prefix, indent string) []byte {
	b, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		common.Panicf("json marshal failed: %v", err)
	}
	return b
}
