package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/glennhartmann/aoc23/src/common/must"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	cardNum := 1
	sum := int64(0)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			log.Printf("EOF")
			break
		}
		if err != nil {
			panic("unable to read")
		}
		s = strings.TrimSuffix(s, "\n")
		log.Printf("current line: %q", s)

		sSplit := strings.Split(s, ": ")
		data := sSplit[1]
		dataSplit := strings.Split(data, " | ")
		winnersStr := strings.Split(dataSplit[0], " ")
		mineStr := strings.Split(dataSplit[1], " ")

		winners := make(map[int]struct{}, len(winnersStr))
		for _, winner := range winnersStr {
			if winner == "" {
				continue
			}
			winners[must.Atoi(winner)] = struct{}{}
		}

		numWinners := 0
		for _, myNum := range mineStr {
			if myNum == "" {
				continue
			}
			num := must.Atoi(myNum)
			if _, ok := winners[num]; ok {
				log.Printf("Card %d: %d is a winning number", cardNum, num)
				numWinners++
			}
		}

		points := int64(0)
		if numWinners > 0 {
			points = intPow(2, numWinners-1)
		}
		sum += points
		log.Printf("Card %d is worth %d points; cumulative sum so far: %d", cardNum, points, sum)

		cardNum++
	}
	log.Printf("total points: %d", sum)
}

func intPow(base, exp int) int64 {
	return int64(math.Pow(float64(base), float64(exp)))
}
