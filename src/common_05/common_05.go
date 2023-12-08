package common_05

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/glennhartmann/aoc23/src/common/must"
)

type RangeMap struct {
	Dst  int64
	Src  int64
	RLen int64
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

func ParseInput() ([]int64, []RangeMap, []RangeMap, []RangeMap, []RangeMap, []RangeMap, []RangeMap, []RangeMap) {
	seeds := make([]int64, 0, 50)
	seedToSoil := make([]RangeMap, 0, 50)
	soilToFertilizer := make([]RangeMap, 0, 50)
	fertilizerToWater := make([]RangeMap, 0, 50)
	waterToLight := make([]RangeMap, 0, 50)
	lightToTemperature := make([]RangeMap, 0, 50)
	temperatureToHumidity := make([]RangeMap, 0, 50)
	humidityToLocation := make([]RangeMap, 0, 50)

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

		rMap := RangeMap{}
		switch state {
		case stateSeeds:
			if s == "" {
				state++
				continue
			}
			seedsSpl := strings.Split(strings.Split(s, ": ")[1], " ")
			for _, ss := range seedsSpl {
				seeds = append(seeds, must.Atoi64(ss))
			}
			log.Printf("seeds: %v", seeds)
			continue
		case stateSeedToSoil, stateSoilToFertilizer, stateFertilizerToWater, stateWaterToLight, stateLightToTemperature, stateTemperatureToHumidity, stateHumidityToLocation:
			if s == "" {
				state++
				continue
			}
			rSpl := strings.Split(s, " ")
			rMap = RangeMap{
				Dst:  must.Atoi64(rSpl[0]),
				Src:  must.Atoi64(rSpl[1]),
				RLen: must.Atoi64(rSpl[2]),
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
			log.Printf("adding seedsToSoil RangeMap: %v", rMap)
		case stateSoilToFertilizer:
			soilToFertilizer = append(soilToFertilizer, rMap)
			log.Printf("adding soilToFertilizer RangeMap: %v", rMap)
		case stateFertilizerToWater:
			fertilizerToWater = append(fertilizerToWater, rMap)
			log.Printf("adding fertilizerToWater RangeMap: %v", rMap)
		case stateWaterToLight:
			waterToLight = append(waterToLight, rMap)
			log.Printf("adding waterToLight RangeMap: %v", rMap)
		case stateLightToTemperature:
			lightToTemperature = append(lightToTemperature, rMap)
			log.Printf("adding lightToTemperature RangeMap: %v", rMap)
		case stateTemperatureToHumidity:
			temperatureToHumidity = append(temperatureToHumidity, rMap)
			log.Printf("adding temperatureToHumidity RangeMap: %v", rMap)
		case stateHumidityToLocation:
			humidityToLocation = append(humidityToLocation, rMap)
			log.Printf("adding humidityToLocation RangeMap: %v", rMap)
		default:
			panic("bad state")
		}
	}
	return seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation
}
