package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"golang.org/x/exp/maps"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	distinct := 4
	if part2 {
		distinct = 14
	}
	last := make([]rune, distinct)
	set := map[rune]bool{}
	unique := func() bool {
		maps.Clear(set)
		for _, r := range last {
			if set[r] {
				return false
			}
			set[r] = true
		}
		return true
	}
	for i, r := range input {
		li := i % distinct
		last[li] = r
		if i >= distinct && unique() {
			// nth character (not zero-indexed)
			return i + 1
		}
	}
	panic("no start of packet found")
}
