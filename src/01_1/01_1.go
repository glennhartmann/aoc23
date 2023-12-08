package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/glennhartmann/aoc23/src/common"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			log.Printf("EOF")
			break
		}
		if err != nil {
			panic("unable to read")
		}
		log.Printf("current line: %q", s)

		first := getFirstDigit(s)
		log.Printf("first digit: %c", first)

		last := getLastDigit(s)
		log.Printf("last digit: %c", last)

		val := getValueForDigits(first, last)
		log.Printf("value: %d", val)

		sum += val
	}
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
