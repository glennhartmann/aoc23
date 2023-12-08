package main

import (
	"log"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/must"
)

func main() {
	sum := 0
	must.ForEachLineOfStreamedInput(func(lineNum int, s string) {
		first := getFirstDigit(s)
		log.Printf("first digit: %c", first)

		last := getLastDigit(s)
		log.Printf("last digit: %c", last)

		val := getValueForDigits(first, last)
		log.Printf("value: %d", val)

		sum += val
	})
	log.Printf("sum: %d", sum)
}

func getFirstDigit(s string) byte {
	for i := 0; i < len(s); i++ {
		if common.IsDigit(s[i]) {
			return s[i]
		}
	}
	panic("no digits found")
}

func getLastDigit(s string) byte {
	for i := len(s) - 1; i >= 0; i-- {
		if common.IsDigit(s[i]) {
			return s[i]
		}
	}
	panic("no digits found")
}

func getValueForDigits(first, last byte) int {
	return common.DigitToInt(first)*10 + common.DigitToInt(last)
}
