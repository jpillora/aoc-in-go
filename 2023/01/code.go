package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			log.Println(input)
		}

		sum += join(firstNum(line, true, part2), firstNum(line, false, part2))
	}
	return sum
}

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func firstNum(line string, front bool, tryWords bool) rune {
	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		c := i
		if !front {
			c = len(runes) - i - 1
		}
		// check for digits
		r := runes[c]
		if r >= '0' && r <= '9' {
			return r
		}
		// check for number words?
		if !tryWords {
			continue
		}
		sub := line[0 : i+1]
		if !front {
			sub = line[c:]
		}
		for i, n := range numbers {
			if strings.Contains(sub, n) {
				return rune(i + '1')
			}
		}
	}
	panic("no number found in: " + line)
}

func join(a, b rune) int {
	n, err := strconv.Atoi(string([]rune{a, b}))
	if err != nil {
		panic(err)
	}
	return n
}
