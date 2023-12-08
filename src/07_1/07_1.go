package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type handBid struct {
	hand     string
	bid      int64
	handType handType
}

const (
	T = 10
	J = 11
	Q = 12
	K = 13
	A = 14
)

type handType int

const (
	highCardType handType = iota
	pairType
	twoPairType
	threeOfAKindType
	fullHouseType
	fourOfAKindType
	fiveOfAKindType
)

func (ht handType) String() string {
	switch ht {
	case highCardType:
		return "highCardType"
	case pairType:
		return "pairType"
	case twoPairType:
		return "twoPairType"
	case threeOfAKindType:
		return "threeOfAKindType"
	case fullHouseType:
		return "fullHouseType"
	case fourOfAKindType:
		return "fourOfAKindType"
	case fiveOfAKindType:
		return "fiveOfAKindType"
	default:
		return "UNKNOWN"
	}
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("error reading from stdin")
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	handBids := make([]handBid, 0, len(lines))
	for i, line := range lines {
		spLine := strings.Split(line, " ")
		handBids = append(handBids, handBid{
			hand:     spLine[0],
			bid:      val(spLine[1]),
			handType: getHandType(spLine[0]),
		})
		log.Printf("line %d: hand %s with bid %s has type %v", i, spLine[0], spLine[1], handBids[len(handBids)-1].handType)
	}

	sort.Slice(handBids, func(i, j int) bool {
		hti := handBids[i].handType
		htj := handBids[j].handType
		if hti > htj {
			return false
		}
		if hti < htj {
			return true
		}
		return cardValuesLess(handBids[i].hand, handBids[j].hand)
	})

	totalWinnings := int64(0)
	for i, handBid := range handBids {
		rank := int64(i + 1)
		winnings := handBid.bid * rank
		log.Printf("(sorted) hand %d (%s) wins %d * %d = %d", i, handBid.hand, handBid.bid, rank, winnings)
		totalWinnings += winnings
	}
	log.Printf("total winnings: %d", totalWinnings)
}

func getHandType(hand string) handType {
	m := make(map[byte]int, len(hand))
	for i := 0; i < len(hand); i++ {
		m[hand[i]]++
	}

	hasFive, hasFour, hasThree, hasTwo := false, false, false, 0
	for _, count := range m {
		switch count {
		case 5:
			hasFive = true
		case 4:
			hasFour = true
		case 3:
			hasThree = true
		case 2:
			hasTwo++
		}
	}

	switch {
	case hasFive:
		return fiveOfAKindType
	case hasFour:
		return fourOfAKindType
	case hasThree && hasTwo > 0:
		return fullHouseType
	case hasThree:
		return threeOfAKindType
	case hasTwo == 2:
		return twoPairType
	case hasTwo == 1:
		return pairType
	default:
		return highCardType
	}
}

func cardValuesLess(h1, h2 string) bool {
	for i := 0; i < len(h1); i++ {
		cv1 := cardVal(h1[i])
		cv2 := cardVal(h2[i])
		if cv1 < cv2 {
			return true
		}
		if cv1 > cv2 {
			return false
		}
	}
	return false
}

func cardVal(c byte) int {
	if isDigit(c) {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return T
	case 'J':
		return J
	case 'Q':
		return Q
	case 'K':
		return K
	case 'A':
		return A
	}
	return -1
}

func val(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("bad strconv")
	}
	return i
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
