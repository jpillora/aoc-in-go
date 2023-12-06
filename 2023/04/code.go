package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
		return sumCopies(input)
	}
	return sumAllWins(input)
}

func sumCopies(input string) int {
	games := strings.Split(input, "\n")
	// make a mirror list of copy counts
	copies := make([]int, len(games))
	for i := range copies {
		copies[i] = 1
	}
	sum := 0
	for i, line := range games {
		c := copies[i]
		w := countWins(line)
		for j := 1; j <= w; j++ {
			copies[i+j] += c
		}
		sum += c
	}
	return sum
}

func sumAllWins(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += sumWins(line)
	}
	return sum
}

func sumWins(line string) int {
	return int(math.Pow(2, float64(countWins(line))))
}

func countWins(line string) int {
	idSets := strings.SplitN(line, ":", 2)
	winHas := strings.SplitN(idSets[1], "|", 2)
	win := set(winHas[0])
	has := set(winHas[1])
	count := 0
	for n := range has {
		if win[n] {
			count++
		}
	}
	return count
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
