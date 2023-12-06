package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	f := forest{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		f = append(f, []tree(line))
	}
	return f.score(!part2)
}

type tree rune

func (t tree) blockedBy(s tree) bool {
	return t <= s
}

type forest [][]tree

func (f forest) width() int {
	return len(f[0])
}

func (f forest) height() int {
	return len(f)
}

func (f forest) score(part1 bool) int {
	score := 0
	for i := range f {
		for j := range f[i] {
			if part1 {
				if f.isVisible(i, j) {
					score++
				}
			} else {
				if s := f.scenicScore(i, j); s > score {
					score = s
				}
			}
		}
	}
	return score
}

func (f forest) inside(i, j int) bool {
	return i >= 0 && j >= 0 && i < f.height() && j < f.width()
}

func (f forest) isVisible(I, J int) bool {
	target := f[I][J]
	for _, d := range dirs {
		visible := true
		for i, j := d.step(I, J); f.inside(i, j); i, j = d.step(i, j) {
			if target.blockedBy(f[i][j]) {
				visible = false // blocked in this direction
				break
			}
		}
		if visible {
			return true
		}
	}
	return false // no line of sight, invisible
}

func (f forest) scenicScore(I, J int) int {
	mult := 1
	target := f[I][J]
	for _, d := range dirs {
		visible := 0
		for i, j := d.step(I, J); f.inside(i, j); i, j = d.step(i, j) {
			visible++
			if target.blockedBy(f[i][j]) {
				break
			}
		}
		if visible == 0 {
			return 0
		}
		mult *= visible
	}
	return mult
}

type dir int

var dirs = []dir{up, down, left, right}

const (
	up dir = iota
	down
	left
	right
)

func (d dir) step(i, j int) (int, int) {
	if d == up {
		i--
	} else if d == down {
		i++
	} else if d == left {
		j--
	} else if d == right {
		j++
	}
	return i, j
}
