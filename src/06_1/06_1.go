package main

import (
	"log"
	"math"
)

type race struct {
	time       int64
	recordDist int64
}

// hardcoded from inputs/06.txt
var input = []race{
	race{
		time:       52,
		recordDist: 426,
	},
	race{
		time:       94,
		recordDist: 1374,
	},
	race{
		time:       75,
		recordDist: 1279,
	},
	race{
		time:       94,
		recordDist: 1216,
	},
}

func main() {
	product := int64(1)
	for i, rc := range input {
		// dist = time_holding_button * (race_time - time_holding_button)
		// => quadratic formula =>
		// time_holding_button = (race_time ± sqrt(race_time^2 - 4 * min_dist)) / 2
		left := (float64(rc.time) - math.Sqrt(float64(rc.time*rc.time)-4*(float64(rc.recordDist)+0.5))) / float64(2)
		right := (float64(rc.time) + math.Sqrt(float64(rc.time*rc.time)-4*(float64(rc.recordDist)+0.5))) / float64(2)
		leftInt := int64(math.Ceil(left))
		rightInt := int64(right)
		factor := rightInt - leftInt + 1
		log.Printf("race %d: [%f -> %d] [%f -> %d] = %d", i, left, leftInt, right, rightInt, factor)
		product *= factor
	}
	log.Printf("product: %v", product)
}
