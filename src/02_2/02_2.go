package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	gameID := 1
	powerSum := 0
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

		s = strings.TrimSpace(s)

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

		gameID++
	}

	log.Printf("sum: %d", powerSum)
}
