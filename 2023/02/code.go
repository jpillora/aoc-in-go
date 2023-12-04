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
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if part2 {
			sum += powerGame(line)
		} else {
			sum += validGameID(line)
		}
	}
	return sum
}

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func validGameID(line string) int {
	id, nc := parseGame(line)
	g := atoi(strings.TrimPrefix(id, "Game "))
	for _, nc := range nc {
		if nc.num > maxCubes[nc.color] {
			return 0 // impossible
		}
	}
	return g
}

func powerGame(line string) int {
	_, nc := parseGame(line)
	// build max values of each color for this game
	max := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, nc := range nc {
		if nc.num > max[nc.color] {
			max[nc.color] = nc.num
		}
	}
	mul := 1
	for _, num := range max {
		mul *= num
	}
	return mul
}

type numColor struct {
	num   int
	color string
}

func parseGame(line string) (id string, nc []numColor) {
	idGames := strings.SplitN(line, ":", 2)
	id = idGames[0]
	for _, game := range strings.Split(idGames[1], ";") {
		for _, cube := range strings.Split(game, ",") {
			pair := strings.SplitN(strings.TrimSpace(cube), " ", 2)
			num, color := atoi(pair[0]), pair[1]
			nc = append(nc, numColor{num, color})
		}
	}
	return
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
