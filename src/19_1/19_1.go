package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/glennhartmann/aoc23/src/common"
	"github.com/glennhartmann/aoc23/src/common/must"
)

func main() {
	lines := must.GetFullInput()
	spLines := common.SplitSlice(lines, []string{""})
	wfs := spLines[0]
	ps := spLines[1]

	workflows := parseWorkflows(wfs)
	log.Printf("workflows: %s", must.JSONMarshalIndent(workflows, "", "  "))

	parts := parseParts(ps)
	log.Printf("parts: %+v", parts)

	sum := 0
	for i, part := range parts {
		nextName := "in"
		path := make([]string, 0, 10)
		path = append(path, nextName)
		for nextName != "A" && nextName != "R" {
			nextName = getNext(part, nextName, workflows)
			path = append(path, nextName)
		}
		log.Printf("%d %+v %+v", i, part, path)
		if nextName == "A" {
			log.Printf("%d %+v accepted", i, part)
			sum += part.x + part.m + part.a + part.s
		} else {
			log.Printf("%d %+v rejected", i, part)
		}
	}
	log.Printf("final sum: %d", sum)
}

func getNext(part part, nextName string, workflows map[string]workflow) string {
	workflow := workflows[nextName]
	for _, c := range workflow.Conditions {
		switch c.CType {
		case ctDeterministic:
			return c.NextName
		case ctGreaterThan:
			if getVariableValue(part, c.Variable) > c.Val {
				return c.NextName
			}
		case ctLessThan:
			if getVariableValue(part, c.Variable) < c.Val {
				return c.NextName
			}
		}
	}
	panic("no conditions match")
}

func getVariableValue(part part, v string) int {
	switch v {
	case "x":
		return part.x
	case "m":
		return part.m
	case "a":
		return part.a
	case "s":
		return part.s
	default:
		panic("bad variable")
	}
}

type workflow struct {
	Name       string        `json:"name"`
	Conditions []wfCondition `json:"conditions"`
}

type conditionType int

const (
	ctDeterministic conditionType = iota
	ctLessThan
	ctGreaterThan
)

func (c conditionType) String() string {
	switch c {
	case ctDeterministic:
		return "deterministic"
	case ctLessThan:
		return "lessThan"
	case ctGreaterThan:
		return "greaterThan"
	default:
		panic("bad condition type")
	}
}

// implements json.Marshaler interface
func (c conditionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.String())), nil
}

type wfCondition struct {
	CType    conditionType `json:"condition_type"`
	Variable string        `json:"variable"`
	Val      int           `json:"value"`
	NextName string        `json:"next_name"`
}

var (
	topLevelWFRx  = regexp.MustCompile(`(\w+){(.*)}`)
	wfConditionRx = regexp.MustCompile(`(\w+)(<|>)(\d+):(\w+)`)
)

func parseWorkflows(wfs []string) map[string]workflow {
	workflows := make(map[string]workflow, len(wfs))
	for _, line := range wfs {
		w := workflow{}
		r := must.FindStringSubmatch(topLevelWFRx, line, 3)
		w.Name = r[1]

		conds := r[2]
		spConds := strings.Split(conds, ",")
		for _, cond := range spConds {
			if strings.Contains(cond, ":") {
				c := wfCondition{}
				r2 := must.FindStringSubmatch(wfConditionRx, cond, 5)
				c.CType = ctLessThan
				if r2[2] == ">" {
					c.CType = ctGreaterThan
				}
				c.Variable = r2[1]
				c.Val = must.Atoi(r2[3])
				c.NextName = r2[4]
				w.Conditions = append(w.Conditions, c)
			} else {
				w.Conditions = append(w.Conditions, wfCondition{CType: ctDeterministic, NextName: cond})
			}
		}

		workflows[w.Name] = w
	}
	return workflows
}

type part struct {
	x, m, a, s int
}

func parseParts(ps []string) []part {
	parts := make([]part, len(ps))
	for i, line := range ps {
		ln := strings.TrimPrefix(strings.TrimSuffix(line, "}"), "{")
		spln := strings.Split(ln, ",")
		getInt := func(i int) int { return must.Atoi(strings.Split(spln[i], "=")[1]) }
		parts[i].x = getInt(0)
		parts[i].m = getInt(1)
		parts[i].a = getInt(2)
		parts[i].s = getInt(3)
	}
	return parts
}
