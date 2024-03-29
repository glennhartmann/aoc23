package main

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/glennhartmann/aoclib/must"
)

type brick struct {
	name       char
	start, end point
}

type char byte

// for spew output
func (c char) String() string {
	return fmt.Sprintf("%c", c)
}

type point struct {
	x, y, z int
}

func main() {
	lines := must.GetFullInput()

	bricks := make([]*brick, 0, len(lines))
	grid := make(map[point]*brick, len(bricks)*10)
	for _, line := range lines {
		spLine := strings.Split(line, "~")
		start := must.ParseListOfNumbers(spLine[0], ",")
		end := must.ParseListOfNumbers(spLine[1], ",")
		b := &brick{
			start: point{
				x: start[0],
				y: start[1],
				z: start[2],
			},
			end: point{
				x: end[0],
				y: end[1],
				z: end[2],
			},
		}
		bricks = append(bricks, b)
		for _, point := range getPoints(b) {
			grid[point] = b
		}
	}

	brickSortFunc := func(a, b *brick) int {
		lowestZA := min(a.start.z, a.end.z)
		lowestZB := min(b.start.z, b.end.z)
		if lowestZA < lowestZB {
			return -1
		}
		if lowestZA > lowestZB {
			return 1
		}
		return 0
	}
	slices.SortFunc(bricks, brickSortFunc)

	ch := char('A')
	for _, b := range bricks {
		b.name = ch
		ch++
		if ch > 'Z' {
			ch = 'A' // doesn't really work after 26, *shrug*
		}
	}

	log.Printf("original bricks:\n%s", spew.Sdump(bricks))

	log.Printf("x:")
	printX(grid)
	log.Printf("")

	log.Printf("y:")
	printY(grid)
	log.Printf("")

	for _, b := range bricks {
		points := getPoints(b)
		for blockedBelow(grid, b, points) == nil {
			moveDown(grid, b, points)
			points = getPoints(b)
		}
	}
	slices.SortFunc(bricks, brickSortFunc)
	log.Printf("bricks after falling:\n%s", spew.Sdump(bricks))

	log.Printf("x:")
	printX(grid)
	log.Printf("")

	log.Printf("y:")
	printY(grid)
	log.Printf("")

	supportingBricks := make(map[*brick][]*brick, len(bricks))
	supportedBricks := make(map[*brick]map[*brick]struct{}, len(bricks))
	for _, b := range bricks {
		supporters := blockedBelow(grid, b, getPoints(b))
		supportedBricks[b] = supporters

		for sb := range supporters {
			supportingBricks[sb] = append(supportingBricks[sb], b)
		}
	}

	for _, b := range bricks {
		log.Printf("%c supports:", b.name)
		for _, sb := range supportingBricks[b] {
			log.Printf("  %c (%d)", sb.name, len(supportedBricks[sb]))
		}
	}

	totalCount := 0
	for _, b := range bricks {
		supported := make(map[*brick]map[*brick]struct{}, len(supportedBricks))
		for k, v := range supportedBricks {
			supported[k] = maps.Clone(v)
		}
		count := disintegrateRecursive(supportingBricks, supported, b)
		log.Printf("disintegrating %c would cause %d other bricks to fall", b.name, count)
		totalCount += count
	}
	log.Printf("total count: %d", totalCount)
}

func disintegrateRecursive(supportingBricks map[*brick][]*brick, supported map[*brick]map[*brick]struct{}, b *brick) int {
	count := 0
	for _, sb := range supportingBricks[b] {
		if len(supported[sb]) == 1 {
			count++
			count += disintegrateRecursive(supportingBricks, supported, sb)
		}
		delete(supported[sb], b)
	}
	return count
}

func getPoints(b *brick) []point {
	maxX, minX := max(b.start.x, b.end.x), min(b.start.x, b.end.x)
	maxY, minY := max(b.start.y, b.end.y), min(b.start.y, b.end.y)
	maxZ, minZ := max(b.start.z, b.end.z), min(b.start.z, b.end.z)

	points := make([]point, 0, 10)
	if maxX != minX {
		for i := minX; i <= maxX; i++ {
			points = append(points, point{
				x: i,
				y: maxY,
				z: maxZ,
			})
		}
	}

	if maxY != minY {
		for i := minY; i <= maxY; i++ {
			points = append(points, point{
				x: maxX,
				y: i,
				z: maxZ,
			})
		}
	}

	if maxZ != minZ {
		for i := minZ; i <= maxZ; i++ {
			points = append(points, point{
				x: maxX,
				y: maxY,
				z: i,
			})
		}
	}

	if maxX == minX && maxY == minY && maxZ == minZ {
		points = append(points, point{
			x: maxX,
			y: maxY,
			z: maxZ,
		})
	}

	return points
}

// nil: not blocked
// map[*brick]struct{}{}: blocked by ground
// map[*brick]struct{}{...}: blocked by bricks
func blockedBelow(grid map[point]*brick, b *brick, points []point) map[*brick]struct{} {
	var b2s map[*brick]struct{}
	for _, p := range points {
		if p.z == 1 {
			return map[*brick]struct{}{}
		}
		if b2, ok := grid[point{
			x: p.x,
			y: p.y,
			z: p.z - 1,
		}]; ok && b2 != b {
			if b2s == nil {
				b2s = make(map[*brick]struct{})
			}
			if _, ok := b2s[b2]; !ok {
				b2s[b2] = struct{}{}
			}
		}
	}
	return b2s
}

func moveDown(grid map[point]*brick, b *brick, points []point) {
	b.start.z--
	b.end.z--

	for _, p := range points {
		delete(grid, p)
	}

	for _, p := range points {
		grid[point{
			x: p.x,
			y: p.y,
			z: p.z - 1,
		}] = b
	}
}

func printX(grid map[point]*brick) {
	printXYInternal(grid,
		func(p point) int { return p.x },
		func(p point) int { return p.y },
		func(r, c, d int) point { return point{x: c, y: d, z: r} })
}

func printY(grid map[point]*brick) {
	printXYInternal(grid,
		func(p point) int { return p.y },
		func(p point) int { return p.x },
		func(r, c, d int) point { return point{x: d, y: c, z: r} })
}

func printXYInternal(grid map[point]*brick, mainVar, depthVar func(p point) int, newPoint func(r, c, d int) point) {
	m := getMax(grid)
	for r := m.z; r > 0; r-- {
		var sb strings.Builder
		for c := 0; c <= mainVar(m); c++ {
			s := char('.')
			for d := 0; d <= depthVar(m); d++ {
				if b, ok := grid[newPoint(r, c, d)]; ok {
					if s != '.' && s != b.name {
						s = '?'
					} else {
						s = b.name
					}
				}
			}
			fmt.Fprintf(&sb, "%c", s)
		}
		log.Printf("%s", sb.String())
	}
}

func getMax(grid map[point]*brick) point {
	m := point{x: -1, y: -1, z: -1}
	for p := range grid {
		m = point{
			x: max(m.x, p.x),
			y: max(m.y, p.y),
			z: max(m.z, p.z),
		}
	}
	return m
}
