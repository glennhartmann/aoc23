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

	minLocation := int64(-1)
	for _, seed := range seeds {
		soil := resolveMaps(seed, seedToSoil)
		fertilizer := resolveMaps(soil, soilToFertilizer)
		water := resolveMaps(fertilizer, fertilizerToWater)
		light := resolveMaps(water, waterToLight)
		temperature := resolveMaps(light, lightToTemperature)
		humidity := resolveMaps(temperature, temperatureToHumidity)
		location := resolveMaps(humidity, humidityToLocation)

		log.Printf("seed %d corresponds to location %d", seed, location)

		if minLocation == int64(-1) || location < minLocation {
			minLocation = location
		}

		log.Printf("min location so far: %d", minLocation)
	}

	log.Printf("min location: %d", minLocation)
}

func resolveMaps(v int64, maps []rangeMap) int64 {
	for _, rMap := range maps {
		if i := findInRangeMap(v, rMap); i != -1 {
			return i
		}
	}
	return v
}

func findInRangeMap(v int64, m rangeMap) int64 {
	if v >= m.src && v < m.src+m.rLen {
		return m.dst + v - m.src
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
