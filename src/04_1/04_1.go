package main

import (
	"log"
	"math"
	"strings"

	"github.com/glennhartmann/aoclib/must"
)

func main() {
	sum := int64(0)
	must.ForEachLineOfStreamedInput(func(lineNum int, s string) {
		cardNum := lineNum + 1
		sSplit := strings.Split(s, ": ")
		data := sSplit[1]
		dataSplit := strings.Split(data, " | ")
		winnersSlice := must.ParseListOfNumbers(dataSplit[0], " ")
		mine := must.ParseListOfNumbers(dataSplit[1], " ")

		winners := make(map[int]struct{}, len(winnersSlice))
		for _, winner := range winnersSlice {
			winners[winner] = struct{}{}
		}

		numWinners := 0
		for _, myNum := range mine {
			if _, ok := winners[myNum]; ok {
				log.Printf("Card %d: %d is a winning number", cardNum, myNum)
				numWinners++
			}
		}

		points := int64(0)
		if numWinners > 0 {
			points = intPow(2, numWinners-1)
		}
		sum += points
		log.Printf("Card %d is worth %d points; cumulative sum so far: %d", cardNum, points, sum)
	})
	log.Printf("total points: %d", sum)
}

func intPow(base, exp int) int64 {
	return int64(math.Pow(float64(base), float64(exp)))
}
