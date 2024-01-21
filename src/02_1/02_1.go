package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/glennhartmann/aoclib/must"
)

func main() {
	idSum := 0
	must.ForEachLineOfStreamedInput(func(lineNum int, s string) {
		gameID := lineNum + 1
		sSplit := strings.Split(s, ": ")
		groups := sSplit[1]
		groupsSplit := strings.Split(groups, "; ")

		log.Printf("Game %d:", gameID)
		couldHappen := true
	outer:
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

				if (colour == "red" && num > 12) || (colour == "green" && num > 13) || (colour == "blue" && num > 14) {
					couldHappen = false
					break outer
				}
			}
		}

		if couldHappen {
			log.Printf("game %d could happen", gameID)
			idSum += gameID
		} else {
			log.Printf("game %d could not happen", gameID)
		}
	})

	log.Printf("sum: %d", idSum)
}
