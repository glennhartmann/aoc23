package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/glennhartmann/aoc23/src/common/must"
)

const (
	//testAreaMin = 7
	//testAreaMax = 27
	testAreaMin = 200000000000000
	testAreaMax = 400000000000000
)

type hail struct {
	xPos, yPos, zPos float64
	xVel, yVel, zVel float64
}

func (h hail) String() string {
	return fmt.Sprintf("%d, %d, %d @ %d, %d, %d", int64(h.xPos), int64(h.yPos), int64(h.zPos), int64(h.xVel), int64(h.yVel), int64(h.zVel))
}

func main() {
	lines := must.GetFullInput()

	sum := 0

	hails := make([]hail, 0, len(lines))
	for _, line := range lines {
		spLine := strings.Split(line, " @ ")
		poss := spLine[0]
		vels := spLine[1]

		spPoss := must.ParseListOfNumbers64(poss, ", ")
		spVels := must.ParseListOfNumbers64(vels, ", ")

		hails = append(hails, hail{
			xPos: float64(spPoss[0]),
			yPos: float64(spPoss[1]),
			zPos: float64(spPoss[2]),
			xVel: float64(spVels[0]),
			yVel: float64(spVels[1]),
			zVel: float64(spVels[2]),
		})
	}

	log.Printf("%+v", hails)

	for i := range hails {
		for j := i + 1; j < len(hails); j++ {
			hailA := hails[i]
			hailB := hails[j]

			log.Printf("Hailstone A: %v", hailA)
			log.Printf("Hailstone B: %v", hailB)

			slopeA := hailA.yVel / hailA.xVel
			slopeB := hailB.yVel / hailB.xVel

			if slopeA == slopeB {
				log.Printf("Hailstones' paths are parallel; they never intersect.\n ")
				continue
			}

			yinterceptA := hailA.yPos - slopeA*hailA.xPos
			yinterceptB := hailB.yPos - slopeB*hailB.xPos

			intersectionX := (yinterceptA - yinterceptB) / (slopeB - slopeA)
			intersectionY := slopeA*intersectionX + yinterceptA

			switch {
			case ((intersectionX < hailA.xPos && hailA.xVel > 0) || (intersectionX > hailA.xPos && hailA.xVel < 0)) && ((intersectionX < hailB.xPos && hailB.xVel > 0) || (intersectionX > hailB.xPos && hailB.xVel < 0)):
				log.Printf("Hailstones' paths crossed in the past for both hailstones.\n ")
			case (intersectionX < hailA.xPos && hailA.xVel > 0) || (intersectionX > hailA.xPos && hailA.xVel < 0):
				log.Printf("Hailstones' paths crossed in the past for hailstone A.\n ")
			case (intersectionX < hailB.xPos && hailB.xVel > 0) || (intersectionX > hailB.xPos && hailB.xVel < 0):
				log.Printf("Hailstones' paths crossed in the past for hailstone B.\n ")
			case intersectionX >= testAreaMin && intersectionX <= testAreaMax && intersectionY >= testAreaMin && intersectionY <= testAreaMax:
				log.Printf("Hailstones' paths will cross inside the test area (at x=%f, y=%f).\n ", intersectionX, intersectionY)
				sum++
			default:
				log.Printf("Hailstones' paths will cross outside the test area (at x=%f, y=%f).\n ", intersectionX, intersectionY)
			}
		}
	}

	log.Printf("%d intersections happened within the test area", sum)
}
