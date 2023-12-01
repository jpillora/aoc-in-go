package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if input == "" {
		return "not implemented"
	}
	count := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		pair := strings.SplitN(line, ",", 2)
		if len(pair) != 2 {
			log.Panicf("invalid input %q", line)
		}
		a := parse(pair[0])
		b := parse(pair[1])
		if !part2 && a.subsetEither(b) {
			count++
		} else if part2 && a.intersetsEither(b) {
			count++
		}
	}
	return count
}

type hilo struct {
	lo, hi int
}

func (h hilo) String() string {
	return fmt.Sprintf("%d-%d", h.lo, h.hi)
}

func parse(s string) hilo {
	pair := strings.SplitN(s, "-", 2)
	if len(pair) != 2 {
		panic("bad pair")
	}
	lo, err := strconv.Atoi(pair[0])
	if err != nil {
		panic(err)
	}
	hi, err := strconv.Atoi(pair[1])
	if err != nil {
		panic(err)
	}
	return hilo{lo, hi}
}

func (a hilo) subsetEither(b hilo) bool {
	return a.subsetOf(b) || b.subsetOf(a)
}

func (a hilo) subsetOf(b hilo) bool {
	return a.lo >= b.lo && a.hi <= b.hi
}

func (a hilo) intersetsEither(b hilo) bool {
	return a.intersets(b) || b.intersets(a)
}

func (a hilo) intersets(b hilo) bool {
	return a.lo <= b.lo && b.lo <= a.hi || // b.lo in a OR
		a.lo <= b.hi && b.hi <= a.hi // b.hi in a
}
