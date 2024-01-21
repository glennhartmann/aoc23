package main

import (
	"log"
	"strings"

	"github.com/glennhartmann/aoc23/src/common/must"

	dll "github.com/glennhartmann/aoclib/doubly_linked_list"
)

type labelVal struct {
	label string
	val   int
}

func main() {
	lines := must.GetFullInput()
	sp := strings.Split(lines[0], ",")

	boxes := make([]*dll.DLL[*labelVal], 256)
	for i := range boxes {
		boxes[i] = dll.NewDLL[*labelVal]()
	}

	for i, step := range sp {
		op := ""
		label := ""
		val := 0
		if strings.HasSuffix(step, "-") {
			op = "-"
			label = strings.TrimSuffix(step, "-")
		} else {
			op = "="
			sp2 := strings.Split(step, "=")
			label = sp2[0]
			val = must.Atoi(sp2[1])
		}

		h := hash(label)
		log.Printf("step %d (%q) hash = %d", i, step, h)

		n := find(boxes[h], label)
		if op == "-" && n != nil {
			log.Printf("removing {%s, %d} from box %d", label, n.Val().val, h)
			n.RemoveFrom(boxes[h])
		} else if op == "=" {
			if n == nil {
				boxes[h].PushTail(&labelVal{label: label, val: val})
				log.Printf("added {%s, %d} to the end of box %d", label, val, h)
			} else {
				n.Val().val = val
				log.Printf("changed %s's val to %d in box %d", label, val, h)
			}
		}
	}

	log.Printf("final state:")
	sum := 0
	for i := range boxes {
		if boxes[i].Len() == 0 {
			continue
		}
		log.Printf("Box %d:", i)
		slot := 1
		for n := boxes[i].Head(); n != nil; n = n.Next() {
			v := (1 + i) * slot * n.Val().val
			log.Printf("  [%s %d] (value %d)", n.Val().label, n.Val().val, v)
			sum += v
			slot++
		}
	}
	log.Printf("final sum: %d", sum)
}

func find(lst *dll.DLL[*labelVal], label string) *dll.Node[*labelVal] {
	for n := lst.Head(); n != nil; n = n.Next() {
		if n.Val().label == label {
			return n
		}
	}
	return nil
}

func hash(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		r = (r + int(s[i])) * 17 % 256
	}
	return r
}
