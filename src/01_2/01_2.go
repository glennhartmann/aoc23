package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
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

		first := getStartOfFirstDigit(s)
		log.Printf("first digit start: %s", first)

		last := getStartOfLastDigit(s)
		log.Printf("last digit start: %s", last)

		val := getValueForDigits(first, last)
		log.Printf("value: %d", val)

		sum += val
	}
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
	return s[0] >= '0' && s[0] <= '9' || isStartOfSpelledOutDigit(s)
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
	if s[0] >= '0' && s[0] <= '9' {
		return int(s[0] - '0')
	}
	return startOfSpelledOutDigitToInt(s)
}
