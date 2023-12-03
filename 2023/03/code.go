package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	rows := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		cols := []rune(line)
		rows = append(rows, cols)
	}

	sum := 0
	for r, row := range rows {
		num := ""
		for c, col := range row {
			if col >= '0' && col <= '9' {
				num += string(col)
			} else if num != "" {
				// found 1 number, check around it for symbols
				if symbolAround(r, c, len(num), rows) {
					sum += atoi(num)
				}
				num = ""
			}
		}
	}
	return sum
}

func symbolAround(r, c int, l int, runes [][]rune) bool {
	// -......-- r,c is location of 7.
	// -.1337.-- and l is 4.
	// -......-- we want to check the dots, not the dashes.
	// IDEA:
	// sweep top l+2 cells
	// sweep left/right two cells
	// sweep bottom l+2 cells
	for y := r - 1; y <= r+1; y++ {
		if y < 0 || y >= len(runes) {
			continue
		}
		for x := c - (l - 1); x <= c+1; x++ {
			if x < 0 || x >= len(runes[y]) {
				continue
			}
			s := symbol(runes[y][x])
			fmt.Printf("%d,%d: %c %v\n", y, x, runes[y][x], s)
			if s {
				return true
			}
		}
	}
	return false
}

func digit(r rune) bool {
	return r >= '0' && r <= '9'
}

func symbol(r rune) bool {
	return !digit(r) && r != '.'
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
