package main

import (
	"log"
	"strings"

	"github.com/glennhartmann/aoc23/src/common/must"

	"github.com/davecgh/go-spew/spew"
	"github.com/glennhartmann/aoc22/src/queue"
)

type moduleType int

const (
	broadcaster moduleType = iota
	flipFlop
	conjunction
	untyped
)

func (mt moduleType) String() string {
	switch mt {
	case broadcaster:
		return "broadcaster"
	case flipFlop:
		return "flipFlop"
	case conjunction:
		return "conjunction"
	case untyped:
		return "untyped"
	default:
		panic("invalid module type")
	}
}

type module struct {
	name        string
	sourceNames []string
	t           moduleType
	destNames   []string
	memory      map[string]pulseType
	ffState     flipFlopState
}

type flipFlopState int

const (
	undefined flipFlopState = iota
	off
	on
)

func (ffs flipFlopState) String() string {
	switch ffs {
	case on:
		return "on"
	case off:
		return "off"
	case undefined:
		return "undefined"
	default:
		panic("invalid flip-flop state")
	}
}

func (ffs flipFlopState) opposite() flipFlopState {
	switch ffs {
	case on:
		return off
	case off:
		return on
	default:
		panic("invalid flip-flop state")
	}
}

type pulseType int

const (
	low pulseType = iota
	high
)

func (pt pulseType) String() string {
	switch pt {
	case high:
		return "high"
	case low:
		return "low"
	default:
		panic("invalid pulse type")
	}
}

type pulse struct {
	sourceName string
	t          pulseType
	destName   string
}

func main() {
	lines := must.GetFullInput()

	modules := make(map[string]*module, len(lines))
	for _, line := range lines {
		spLine := strings.Split(line, " -> ")
		mNameType := spLine[0]
		dests := spLine[1]
		spDests := strings.Split(dests, ", ")

		name := "broadcaster"
		if mNameType == "broadcaster" {
			if modules[name] == nil {
				modules[name] = &module{}
			}
			modules[name].name = mNameType
			modules[name].t = broadcaster
			modules[name].destNames = spDests
		} else if strings.HasPrefix(mNameType, "%") {
			name = strings.TrimPrefix(mNameType, "%")
			if modules[name] == nil {
				modules[name] = &module{}
			}
			modules[name].name = name
			modules[name].t = flipFlop
			modules[name].destNames = spDests
			modules[name].ffState = off
		} else if strings.HasPrefix(mNameType, "&") {
			name = strings.TrimPrefix(mNameType, "&")
			if modules[name] == nil {
				modules[name] = &module{}
			}
			modules[name].name = name
			modules[name].t = conjunction
			modules[name].destNames = spDests
		}

		for _, dn := range modules[name].destNames {
			if modules[dn] == nil {
				modules[dn] = &module{t: untyped}
			}
			modules[dn].sourceNames = append(modules[dn].sourceNames, name)
		}
	}

	for _, module := range modules {
		if module.t != conjunction {
			continue
		}

		module.memory = make(map[string]pulseType)
		for _, sn := range module.sourceNames {
			module.memory[sn] = low
		}
	}

	log.Printf("modules: %s", spew.Sdump(modules))

	lowCount, highCount := 0, 0

	pulses := queue.NewQueue[pulse]()
	for i := 0; i < 1000; i++ {
		pulses.Push(pulse{
			sourceName: "button",
			t:          low,
			destName:   "broadcaster",
		})
		for !pulses.Empty() {
			pls, _ := pulses.Pop()
			dest := modules[pls.destName]

			if pls.t == low {
				lowCount++
			} else {
				highCount++
			}

			log.Printf("%s -%v-> %s", pls.sourceName, pls.t, pls.destName)

			newPulseType := low
			sendPulse := true
			switch dest.t {
			case broadcaster:
				newPulseType = pls.t
			case flipFlop:
				if pls.t == low {
					dest.ffState = dest.ffState.opposite()
					if dest.ffState == on {
						newPulseType = high
					}
				} else {
					sendPulse = false
				}
			case conjunction:
				dest.memory[pls.sourceName] = pls.t
				if !allHigh(dest.memory) {
					newPulseType = high
				}
			case untyped:
				continue
			default:
				panic("bad module type")
			}

			if sendPulse {
				for _, dn := range dest.destNames {
					pulses.Push(pulse{
						sourceName: dest.name,
						t:          newPulseType,
						destName:   dn,
					})
				}
			}
		}
	}

	log.Printf("%d high and %d low -> %d", highCount, lowCount, lowCount*highCount)
}

func allHigh(memory map[string]pulseType) bool {
	for _, pt := range memory {
		if pt == low {
			return false
		}
	}
	return true
}
