package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type rangeMap struct {
	dst  int64
	src  int64
	rLen int64
}

type normalRange struct {
	start int64
	rLen  int64
}

type stateType int

const (
	stateSeeds stateType = iota
	stateSeedToSoilHeader
	stateSeedToSoil
	stateSoilToFertilizerHeader
	stateSoilToFertilizer
	stateFertilizerToWaterHeader
	stateFertilizerToWater
	stateWaterToLightHeader
	stateWaterToLight
	stateLightToTemperatureHeader
	stateLightToTemperature
	stateTemperatureToHumidityHeader
	stateTemperatureToHumidity
	stateHumidityToLocationHeader
	stateHumidityToLocation
	stateEnd
)

func main() {
	seeds := make([]int64, 0, 50)
	seedToSoil := make([]rangeMap, 0, 50)
	soilToFertilizer := make([]rangeMap, 0, 50)
	fertilizerToWater := make([]rangeMap, 0, 50)
	waterToLight := make([]rangeMap, 0, 50)
	lightToTemperature := make([]rangeMap, 0, 50)
	temperatureToHumidity := make([]rangeMap, 0, 50)
	humidityToLocation := make([]rangeMap, 0, 50)

	state := stateSeeds

	r := bufio.NewReader(os.Stdin)
outer:
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

		rMap := rangeMap{}
		switch state {
		case stateSeeds:
			if s == "" {
				state++
				continue
			}
			seedsSpl := strings.Split(strings.Split(s, ": ")[1], " ")
			for _, ss := range seedsSpl {
				seeds = append(seeds, val(ss))
			}
			log.Printf("seeds: %v", seeds)
			continue
		case stateSeedToSoil, stateSoilToFertilizer, stateFertilizerToWater, stateWaterToLight, stateLightToTemperature, stateTemperatureToHumidity, stateHumidityToLocation:
			if s == "" {
				state++
				continue
			}
			rSpl := strings.Split(s, " ")
			rMap = rangeMap{
				dst:  val(rSpl[0]),
				src:  val(rSpl[1]),
				rLen: val(rSpl[2]),
			}
		case stateSeedToSoilHeader, stateSoilToFertilizerHeader, stateFertilizerToWaterHeader, stateWaterToLightHeader, stateLightToTemperatureHeader, stateTemperatureToHumidityHeader, stateHumidityToLocationHeader:
			state++
			continue
		case stateEnd:
			break outer
		default:
			panic("bad state")
		}

		switch state {
		case stateSeedToSoil:
			seedToSoil = append(seedToSoil, rMap)
			log.Printf("adding seedsToSoil rangeMap: %v", rMap)
		case stateSoilToFertilizer:
			soilToFertilizer = append(soilToFertilizer, rMap)
			log.Printf("adding soilToFertilizer rangeMap: %v", rMap)
		case stateFertilizerToWater:
			fertilizerToWater = append(fertilizerToWater, rMap)
			log.Printf("adding fertilizerToWater rangeMap: %v", rMap)
		case stateWaterToLight:
			waterToLight = append(waterToLight, rMap)
			log.Printf("adding waterToLight rangeMap: %v", rMap)
		case stateLightToTemperature:
			lightToTemperature = append(lightToTemperature, rMap)
			log.Printf("adding lightToTemperature rangeMap: %v", rMap)
		case stateTemperatureToHumidity:
			temperatureToHumidity = append(temperatureToHumidity, rMap)
			log.Printf("adding temperatureToHumidity rangeMap: %v", rMap)
		case stateHumidityToLocation:
			humidityToLocation = append(humidityToLocation, rMap)
			log.Printf("adding humidityToLocation rangeMap: %v", rMap)
		default:
			panic("bad state")
		}
	}

	seedRanges := make([]normalRange, 0, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, normalRange{
			start: seeds[i],
			rLen:  seeds[i+1],
		})
		log.Printf("adding normalRange of seeds: %v", seedRanges[len(seedRanges)-1])
	}

	soil := resolveMaps(seedRanges, seedToSoil)
	fertilizer := resolveMaps(soil, soilToFertilizer)
	water := resolveMaps(fertilizer, fertilizerToWater)
	light := resolveMaps(water, waterToLight)
	temperature := resolveMaps(light, lightToTemperature)
	humidity := resolveMaps(temperature, temperatureToHumidity)
	location := resolveMaps(humidity, humidityToLocation)

	log.Printf("seed ranges %v correspond to location ranges %v", seedRanges, location)

	minLocation := int64(-1)
	for _, lcn := range location {
		if minLocation == int64(-1) || lcn.start < minLocation {
			minLocation = lcn.start
		}
	}

	log.Printf("min location: %d", minLocation)
}

var badRange = normalRange{-1, -1}

func resolveMaps(v []normalRange, maps []rangeMap) []normalRange {
	ret := make([]normalRange, 0, 500)
outer:
	for i := 0; i < len(v); i++ {
		rng := v[i]
		for _, rMap := range maps {
			left, match, right := findInRangeMap(rng, rMap)
			if match == badRange {
				continue
			}

			matchCpy := match
			match.start = rMap.dst + match.start - rMap.src
			log.Printf("converting match from src space (%v) to dst space (%v)", matchCpy, match)

			ret = append(ret, match)
			if left != badRange {
				v = append(v, left)
			}
			if right != badRange {
				v = append(v, right)
			}
			continue outer
		}
		ret = append(ret, rng)
	}
	return ret
}

func findInRangeMap(v normalRange, m rangeMap) (left, match, right normalRange) {
	match = getOverlap(v, m)
	if match == badRange {
		return badRange, badRange, badRange
	}

	left, right = badRange, badRange

	if match.start > v.start {
		left = normalRange{
			start: v.start,
			rLen:  match.start - v.start,
		}
	}

	matchEnd := match.start + match.rLen - 1
	vEnd := v.start + v.rLen - 1
	if matchEnd < vEnd {
		right = normalRange{
			start: matchEnd + 1,
			rLen:  vEnd - matchEnd,
		}
	}

	log.Printf("splitting {%v; %v} into {%v; %v; %v}", v, m, left, match, right)
	return left, match, right
}

func getOverlap(v normalRange, m rangeMap) normalRange {
	mEnd := m.src + m.rLen - 1
	vEnd := v.start + v.rLen - 1
	if v.start > mEnd || vEnd < m.src {
		return badRange
	}

	start := max(v.start, m.src)
	return normalRange{
		start: start,
		rLen:  min(vEnd, mEnd) - start + 1,
	}
}

func val(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("bad strconv")
	}
	return i
}
