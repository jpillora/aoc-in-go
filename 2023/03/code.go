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

func run(part2 bool, input string) any {
	runes := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		runes = append(runes, []rune(line))
	}
	if part2 {
		return sumGearRatios(runes)
	}
	return sumSymbolNumbers(runes)
}

func sumGearRatios(runes [][]rune) int {
	sum := 0
	for r, row := range runes {
		for c, col := range row {
			// TODO: find 2 SEPARATE numbers around r/c
			nums := []string{}
			if col == '*' {
				around(r, c, string(col), runes, func(a rune) {
					fmt.Printf("star[%d,%d] -> %s\n", r, c, string(a))
					if digit(a) {
						nums = append(nums, numAt(r, c, runes))
					}
				})
			}
			if len(nums) > 0 {
				fmt.Printf("star[%d,%d] nums %v\n", r, c, nums)
			}
			if len(nums) == 2 {
				sum += atoi(nums[0]) * atoi(nums[1])
			}
		}
	}
	return sum
}

func sumSymbolNumbers(runes [][]rune) int {
	sum := 0
	for r, row := range runes {
		num := ""
		for c, col := range row {
			if digit(col) {
				num += string(col)
			}
			eon := digit(col) && (c == len(row)-1 || !digit(row[c+1]))
			if eon {
				if symbolAround(r, c, num, runes) {
					sum += atoi(num)
				}
				num = ""
			}
		}
	}
	return sum
}

func symbolAround(r, c int, target string, runes [][]rune) bool {
	s := false
	around(r, c, target, runes, func(r rune) {
		if symbol(r) {
			s = true
		}
	})
	return s
}

func numAt(r, c int, runes [][]rune) string {
	// TODO: need location of string?
	row := runes[r]
	num := []rune{
		row[c],
	}
	// check left
	for x := c - 1; x >= 0 && digit(row[x]); x-- {
		num = append([]rune{row[x]}, num...)
	}
	// check right
	for x := c + 1; x < len(row) && digit(row[x]); x++ {
		num = append(num, row[x])
	}
	return string(num)
}

// -......-- r,c is location of the "7".
// -.1337.-- l(ength) of target is 4.
// -......-- we want to check around the target.
func around(r, c int, target string, runes [][]rune, handle func(r rune)) {
	l := len(target)
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
			handle(runes[y][x])
		}
	}
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
