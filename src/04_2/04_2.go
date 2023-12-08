package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/glennhartmann/aoc23/src/common/must"
)

const cardNums = 204 // hardcoded based on inputs/04.txt

func main() {
	r := bufio.NewReader(os.Stdin)
	cardNum := 1

	cardInstances := make([]int64, cardNums+1)
	for i := 1; i < len(cardInstances); i++ {
		cardInstances[i] = int64(1)
	}

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

		for i := 0; i < numWinners; i++ {
			currCard := cardNum + i + 1
			cardInstances[currCard] += cardInstances[cardNum]
			log.Printf("Card %d: adding %d more instances of card %d for a total of %d instances of card %d", cardNum, cardInstances[cardNum], currCard, cardInstances[currCard], currCard)
		}

		cardNum++
	}

	sum := int64(0)
	for i := 1; i < len(cardInstances); i++ {
		sum += cardInstances[i]
	}
	log.Printf("total cards: %d", sum)
}
