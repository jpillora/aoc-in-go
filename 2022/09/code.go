package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

const drawAll = false

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	notches := 2
	if part2 {
		notches = 10
	}
	initial := position{0, 0}
	w := world{
		initial: initial,
		head:    newJourney('H', initial),
		body:    make([]*notch, notches-1), // head included
	}
	for b := 0; b < len(w.body); b++ {
		w.body[b] = newJourney(rune('1'+b), initial)
	}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		m := move{}
		m.unmarshal(line)
		w.simulate(m)
	}
	if len(lines) < 20 {
		w.drawTailNotch()
	}
	return len(w.tail().visited)
}

type world struct {
	initial position
	head    *notch
	body    []*notch
}

func (w *world) notches() []*notch {
	return append([]*notch{w.head}, w.body...)
}

func (w *world) simulate(m move) {
	for c := 1; c <= m.count; c++ {
		// move head notch
		w.head.step(m.dir)
		prev := w.head.position()
		// move body notches
		for _, b := range w.body {
			// if prev notch not adjacent, move
			curr := b.position()
			if !prev.adjacent(curr) {
				next := curr.delta(prev)
				b.next(next)
			}
			prev = b.position()
		}
		if drawAll {
			w.drawAll()
		}
	}
}

func (w *world) tail() *notch {
	return w.body[len(w.body)-1]
}

func (w *world) drawAll() {
	fmt.Println("all notches:")
	minp := w.initial
	maxp := w.initial
	for _, n := range w.notches() {
		p := n.position()
		minp.x = min(minp.x, p.x)
		minp.y = min(minp.y, p.y)
		maxp.x = max(maxp.x, p.x)
		maxp.y = max(maxp.y, p.y)
	}
	for j := minp.y - 1; j <= maxp.y+1; j++ {
		for i := minp.x - 1; i <= maxp.x+1; i++ {
			r := rune('·')
			cp := position{i, j}
			if cp == w.initial {
				r = rune('s')
			}
			for _, n := range w.notches() {
				np := n.position()
				if cp == np {
					r = n.name
					break
				}
			}
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func (w *world) drawTailNotch() {
	minp := w.initial
	maxp := w.initial
	for _, p := range w.tail().steps {
		minp.x = min(minp.x, p.x)
		minp.y = min(minp.y, p.y)
		maxp.x = max(maxp.x, p.x)
		maxp.y = max(maxp.y, p.y)
	}
	for j := minp.y - 1; j <= maxp.y+1; j++ {
		for i := minp.x - 1; i <= maxp.x+1; i++ {
			p := position{i, j}
			n := w.tail().visited[p]
			if n > 0 {
				fmt.Printf("#")
				// fmt.Printf("%d", n)
			} else {
				fmt.Print("·")
			}
		}
		fmt.Println()
	}
}

type notch struct {
	name    rune
	initial position
	steps   []position
	visited map[position]int
}

func newJourney(name rune, initial position) *notch {
	return &notch{
		name:    name,
		initial: initial,
		visited: map[position]int{
			initial: 1,
		},
	}
}

func (j *notch) position() position {
	if len(j.steps) == 0 {
		return j.initial
	}
	return j.steps[len(j.steps)-1]
}

func (j *notch) step(d dir) {
	j.next(j.position().shift(d))
}

func (j *notch) next(p position) {
	j.steps = append(j.steps, p)
	j.visited[p]++
}

type position struct {
	x, y int
}

func (p position) shift(dir dir) position {
	switch dir {
	case up:
		p.y--
	case down:
		p.y++
	case left:
		p.x--
	case right:
		p.x++
	}
	return p
}

// delta position 1 step from p to q
func (p position) delta(q position) position {
	d := p
	// if aligned? move in 1 dimension
	// else (diagonal)? move in 2 dimensions
	if p.x != q.x {
		if p.x < q.x {
			d.x++
		} else {
			d.x--
		}
	}
	if p.y != q.y {
		if p.y < q.y {
			d.y++
		} else {
			d.y--
		}
	}
	return d
}

func (p position) adjacent(q position) bool {
	xd := math.Abs(float64(p.x) - float64(q.x))
	yd := math.Abs(float64(p.y) - float64(q.y))
	return xd <= 1 && yd <= 1
}

// move holds direction, and count. e.g. "U 2"
type move struct {
	dir   dir
	count int
}

func (m *move) unmarshal(s string) {
	re := regexp.MustCompile(`^([UDLR]) (\d+)$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		log.Panicf("invalid move: '%s'", s)
	}
	m.dir = dir(matches[1][0])
	m.count, _ = strconv.Atoi(matches[2])
}

type dir rune

const (
	up    dir = 'U'
	down  dir = 'D'
	left  dir = 'L'
	right dir = 'R'
)

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
