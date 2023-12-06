package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	c := cave{
		sandSource: position{500, 0},
	}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		r := rock{}
		for _, p := range strings.Split(line, " -> ") {
			xy := strings.SplitN(p, ",", 2)
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			r.parts = append(r.parts, position{x, y})
		}
		c.rocks = append(c.rocks, r)
	}
	if part2 {
		c.hasFloor = true
		c.floorY = c.lowestRock() + 2
	}
	w := world{
		cave:  c,
		grain: nil,
		sand:  []*grain{},
	}
	for w.tick() {
		// simulating...
	}
	const debug = false
	if debug {
		w.print()
	}
	return w.resting()
}

type world struct {
	cave
	grain *grain
	sand  []*grain
}

func (w *world) print() {
	minp := position{math.MaxInt, math.MaxInt}
	maxp := position{}
	for _, p := range w.positions() {
		minp.x = min(minp.x, p.x)
		minp.y = min(minp.y, p.y)
		maxp.x = max(maxp.x, p.x)
		maxp.y = max(maxp.y, p.y)
	}
	for j := minp.y - 1; j <= maxp.y+1; j++ {
		for i := minp.x - 1; i <= maxp.x+1; i++ {
			p := position{i, j}
			if w.sandAt(p) {
				fmt.Print("o")
			} else if p == w.sandSource {
				fmt.Print("+")
			} else if j == w.sandSource.y {
				fmt.Print("-")
			} else if w.rockAt(p) {
				fmt.Print("#")
			} else if w.floorAt(p) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (w *world) positions() []position {
	ps := []position{w.sandSource}
	for _, g := range w.cave.rocks {
		ps = append(ps, g.parts...)
	}
	for _, g := range w.sand {
		ps = append(ps, g.position)
	}
	return ps
}

func (w *world) tick() bool {
	g := w.grain
	if g == nil || g.resting {
		g = &grain{position: w.sandSource}
		w.sand = append(w.sand, g)
		w.grain = g
	}
	// move grain
	next := position{}
	canMove := false
	for _, candidate := range g.candidates() {
		if !w.blockedAt(candidate) {
			canMove = true
			next = candidate
			break
		}
	}
	// cant move, mark resting
	if !canMove {
		g.resting = true
		// has floor? sand wont freefall. instead check if cant move from the source
		if w.cave.hasFloor && g.position == w.sandSource {
			w.grain = nil
			return false
		}
		return true
	}
	// cant move
	g.position = next
	g.moves++
	// past lowest rock? freefalling
	if !w.hasFloor {
		freefall := g.position.y > w.cave.lowestRock()
		return !freefall
	}
	return true
}

func (w *world) resting() int {
	r := 0
	for _, g := range w.sand {
		if g.resting {
			r++
		}
	}
	return r
}

func (w *world) blockedAt(p position) bool {
	return w.rockAt(p) || w.sandAt(p) || w.floorAt(p)
}

func (w *world) sandAt(p position) bool {
	for _, g := range w.sand {
		if g == w.grain {
			continue
		}
		if !g.resting {
			panic("all encountered sand should be resting")
		}
		if g.position == p {
			return true
		}
	}
	return false
}

type grain struct {
	position
	moves   int
	resting bool
}

func (g *grain) candidates() []position {
	return []position{
		{g.x, g.y + 1},
		{g.x - 1, g.y + 1},
		{g.x + 1, g.y + 1},
	}
}

type cave struct {
	sandSource position
	rocks      []rock
	hasFloor   bool
	floorY     int
}

func (c *cave) rockAt(p position) bool {
	for _, r := range c.rocks {
		if r.contains(p) {
			return true
		}
	}
	return false
}

func (c *cave) floorAt(p position) bool {
	return c.hasFloor && p.y == c.floorY
}

func (c *cave) lowestRock() int {
	maxY := 0
	for _, r := range c.rocks {
		for _, p := range r.parts {
			if p.y > maxY {
				maxY = p.y
			}
		}
	}
	return maxY
}

type rock struct {
	parts []position
}

func (r *rock) contains(p position) bool {
	from := r.parts[0]
	for _, to := range r.parts[1:] {
		// either x or y will be 1 dimensional
		x := from.x <= p.x && p.x <= to.x || from.x >= p.x && p.x >= to.x
		y := from.y <= p.y && p.y <= to.y || from.y >= p.y && p.y >= to.y
		c := x && y
		// fmt.Printf("contains %v < %v < %v = %v\n", from, p, to, c)
		if c {
			return true
		}
		from = to
	}
	return false
}

type position struct {
	x, y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
