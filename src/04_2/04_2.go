package main

import (
	"log"
	"strings"

	"github.com/glennhartmann/aoc23/src/common/must"
)

const cardNums = 204 // hardcoded based on inputs/04.txt

func main() {
	cardInstances := make([]int64, cardNums+1)
	for i := 1; i < len(cardInstances); i++ {
		cardInstances[i] = int64(1)
	}

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

		for i := 0; i < numWinners; i++ {
			currCard := cardNum + i + 1
			cardInstances[currCard] += cardInstances[cardNum]
			log.Printf("Card %d: adding %d more instances of card %d for a total of %d instances of card %d", cardNum, cardInstances[cardNum], currCard, cardInstances[currCard], currCard)
		}
	})

	sum := int64(0)
	for i := 1; i < len(cardInstances); i++ {
		sum += cardInstances[i]
	}
	log.Printf("total cards: %d", sum)
}
