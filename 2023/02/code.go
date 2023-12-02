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
	idGames := strings.SplitN(line, ":", 2)
	id, games := idGames[0], idGames[1]
	g := atoi(strings.TrimPrefix(id, "Game "))
	for _, game := range strings.Split(games, ";") {
		for _, cube := range strings.Split(game, ",") {
			numColor := strings.SplitN(strings.TrimSpace(cube), " ", 2)
			num, color := atoi(numColor[0]), numColor[1]
			if num > maxCubes[color] {
				return 0 // impossible
			}
		}
	}
	return g
}

func powerGame(line string) int {
	idGames := strings.SplitN(line, ":", 2)
	max := map[string]int{"red": 0, "green": 0, "blue": 0}
	// build max values of each color for this game
	for _, game := range strings.Split(idGames[1], ";") {
		for _, cube := range strings.Split(game, ",") {
			numColor := strings.SplitN(strings.TrimSpace(cube), " ", 2)
			num, color := atoi(numColor[0]), numColor[1]
			if num > max[color] {
				max[color] = num
			}
		}
	}
	mul := 1
	for _, num := range max {
		mul *= num
	}
	return mul
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
