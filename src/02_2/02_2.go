package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/glennhartmann/aoclib/must"
)

func main() {
	powerSum := 0
	must.ForEachLineOfStreamedInput(func(lineNum int, s string) {
		gameID := lineNum + 1
		sSplit := strings.Split(s, ": ")
		groups := sSplit[1]
		groupsSplit := strings.Split(groups, "; ")

		maxGroups := map[string]int{"red": 0, "blue": 0, "green": 0}

		log.Printf("Game %d:", gameID)
		for _, group := range groupsSplit {
			log.Printf("  group:")
			groupSplit := strings.Split(group, ", ")
			for _, cubes := range groupSplit {
				numCol := strings.Split(cubes, " ")
				num, err := strconv.Atoi(numCol[0])
				if err != nil {
					panic("bad number")
				}
				colour := numCol[1]

				log.Printf("    num: %d, colour: %s", num, colour)

				if num > maxGroups[colour] {
					maxGroups[colour] = num
				}
			}
		}

		powerSum += maxGroups["red"] * maxGroups["green"] * maxGroups["blue"]
	})

	log.Printf("sum: %d", powerSum)
}
