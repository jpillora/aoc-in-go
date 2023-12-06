package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += sumGame(line)
	}
	return sum
}

func sumGame(line string) int {
	idSets := strings.SplitN(line, ":", 2)
	winHas := strings.SplitN(idSets[1], "|", 2)
	win := set(winHas[0])
	has := set(winHas[1])
	sum := 0
	for n := range has {
		if win[n] {
			if sum == 0 {
				sum = 1
			} else {
				sum *= 2
			}
		}
	}
	return sum
}

func set(s string) map[int]bool {
	m := map[int]bool{}
	for _, n := range strings.Split(s, " ") {
		ns := strings.TrimSpace(n)
		if ns == "" {
			continue
		}
		m[atoi(ns)] = true
	}
	return m
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
