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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	rows := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		rows = append(rows, []rune(line))
	}
	sum := int64(0)
	for r, row := range rows {
		num := ""
		for c, col := range row {
			if digit(col) {
				num += string(col)
				continue
			}
			if num == "" {
				continue
			}
			if symbolAround(r, c-1, num, rows) {
				sum += int64(atoi(num))
			}
			num = ""
		}
		if num != "" && symbolAround(r, len(row)-1, num, rows) {
			sum += int64(atoi(num))
		}
	}
	return sum
}

// -......-- r,c is location of 7.
// -.1337.-- and l(ength) is 4.
// -......-- we want to check the dots, not the dashes.
func symbolAround(r, c int, num string, runes [][]rune) bool {
	l := len(num)
	y0 := r - 1
	for y := y0; y <= r+1; y++ {
		if y < 0 || y >= len(runes) {
			continue
		}
		x0 := c - l
		for x := x0; x <= c+1; x++ {
			if y == y0+1 && x == x0+1 {
				x += l
			}
			if x < 0 || x >= len(runes[y]) {
				continue
			}
			s := symbol(runes[y][x])
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
