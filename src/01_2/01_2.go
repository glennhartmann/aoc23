package main

import (
	"log"
	"strings"

	"github.com/glennhartmann/aoclib/common"
	"github.com/glennhartmann/aoclib/must"
)

func main() {
	sum := 0
	must.ForEachLineOfStreamedInput(func(lineNum int, s string) {
		first := getStartOfFirstDigit(s)
		log.Printf("first digit start: %s", first)

		last := getStartOfLastDigit(s)
		log.Printf("last digit start: %s", last)

		val := getValueForDigits(first, last)
		log.Printf("value: %d", val)

		sum += val
	})
	log.Printf("sum: %d", sum)
}

func getStartOfFirstDigit(s string) string {
	for i := 0; i < len(s); i++ {
		if isStartOfDigit(s[i:]) {
			return s[i:]
		}
	}
	panic("no digits found")
}

func getStartOfLastDigit(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if isStartOfDigit(s[i:]) {
			return s[i:]
		}
	}
	panic("no digits found")
}

func isStartOfDigit(s string) bool {
	return common.IsDigit(s[0]) || isStartOfSpelledOutDigit(s)
}

func isStartOfSpelledOutDigit(s string) bool {
	return startOfSpelledOutDigitToInt(s) != -1
}

func startOfSpelledOutDigitToInt(s string) int {
	switch {
	case strings.HasPrefix(s, "one"):
		return 1
	case strings.HasPrefix(s, "two"):
		return 2
	case strings.HasPrefix(s, "three"):
		return 3
	case strings.HasPrefix(s, "four"):
		return 4
	case strings.HasPrefix(s, "five"):
		return 5
	case strings.HasPrefix(s, "six"):
		return 6
	case strings.HasPrefix(s, "seven"):
		return 7
	case strings.HasPrefix(s, "eight"):
		return 8
	case strings.HasPrefix(s, "nine"):
		return 9
	}
	return -1
}

func getValueForDigits(first, last string) int {
	return startOfDigitToInt(first)*10 + startOfDigitToInt(last)
}

func startOfDigitToInt(s string) int {
	if common.IsDigit(s[0]) {
		return common.DigitToInt(s[0])
	}
	return startOfSpelledOutDigitToInt(s)
}
