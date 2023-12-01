package main

import (
	"bufio"
	"io"
	"log"
	"os"
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
		if isDigit(s[i]) {
			return s[i]
		}
	}
	panic("no digits found")
}

func getLastDigit(s string) byte {
	for i := len(s) - 1; i >= 0; i-- {
		if isDigit(s[i]) {
			return s[i]
		}
	}
	panic("no digits found")
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func getValueForDigits(first, last byte) int {
	return digitToInt(first)*10 + digitToInt(last)
}

func digitToInt(b byte) int {
	return int(b - '0')
}
