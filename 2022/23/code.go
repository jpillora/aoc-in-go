package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 || len(input) > 1000 {
		return nil
	}
	w := world{}
	for x, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for y, r := range line {
			if r == '#' {
				e := &elf{
					position: position{x, y},
				}
				w.elves = append(w.elves, e)
			}
		}
	}

	w.tick()
	// some elves will have a next position
	for i, e := range w.elves {
		fmt.Printf("elf %d: %v -> %v\n", i, e.position, e.nextPos)
	}
	return len(w.elves)
}

type world struct {
	elves []*elf
}

func (w *world) occupied(p position) bool {
	for _, e := range w.elves {
		if e.position == p {
			return true
		}
	}
	return false
}

func (w *world) tick() bool {
	// first half
	for _, e := range w.elves {
		alone := true
		for _, p := range e.position.around() {
			for _, ee := range w.elves {
				if ee.position == p {
					// found another elf
					alone = false
					goto outer
				}
			}
		}
	outer:
		if alone {
			continue // do nothing
		}
		// propose move
		for _, d := range []dir{N, S, W, E} {
			available := true
			for _, dd := range e.position.fanout(d) {
				if w.occupied(dd) {
					available = false
					break
				}
			}
			if !available {
				continue
			}
			n := d.adjacent(e.position)
			e.nextPos = &n
			break
		}
	}
	// second half, move all elves which do not collide
	// TODO
	return true
}

type elf struct {
	position
	nextPos *position
}

type position struct {
	x, y int
}

func (p position) around() [8]position {
	return [8]position{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

func (p position) fanout(d dir) [3]position {
	adj := d.adjacent(p)
	if adj.x == p.x {
		return [3]position{
			{x: adj.x, y: adj.y - 1},
			adj,
			{x: adj.x, y: adj.y + 1},
		}
	}
	return [3]position{
		{x: adj.x - 1, y: adj.y},
		adj,
		{x: adj.x + 1, y: adj.y},
	}
}

type dir int

const (
	N dir = iota + 1
	S
	W
	E
)

func (d dir) adjacent(p position) position {
	switch d {
	case N:
		return position{p.x - 1, p.y}
	case S:
		return position{p.x + 1, p.y}
	case W:
		return position{p.x, p.y - 1}
	case E:
		return position{p.x, p.y + 1}
	}
	panic("invalid direction")
}
