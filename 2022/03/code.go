package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rs := make(rucksacks, len(lines))
	for i, line := range lines {
		rs[i] = rucksack(line)
	}
	out, err := rs.priority(!part2)
	if err != nil {
		panic(err)
	}
	return out
}

type item rune

func (i item) priority() int {
	if i >= 'a' && i <= 'z' {
		return 1 + int(i-'a')
	}
	if i >= 'A' && i <= 'Z' {
		return 27 + int(i-'A')
	}
	panic("invalid item")
}

type groups []*group

const groupSize = 3

type group [groupSize]rucksack

func (g group) String() string {
	return string(g[0]) + "/" + string(g[1]) + "/" + string(g[2])
}

type rucksacks []rucksack

func (rs rucksacks) priority(dupe bool) (int, error) {
	if dupe {
		return rs.dupePriority()
	}
	return rs.groupPriority()
}

func (rs rucksacks) groupPriority() (int, error) {
	sum := 0
	for _, g := range rs.split() {
		triplicate, err := g[0].intersect(g[1]).intersect(g[2]).one()
		if err != nil {
			return 0, fmt.Errorf("group %s: %s", g, err)
		}
		sum += triplicate.priority()
	}
	return sum, nil
}

func (rs rucksacks) dupePriority() (int, error) {
	sum := 0
	for _, r := range rs {
		left, right := r.split()
		dupe, err := left.intersect(right).one()
		if err != nil {
			return sum, fmt.Errorf("left %s right %s: %s", string(left), string(right), err)
		}
		sum += dupe.priority()
	}
	return sum, nil
}

func (rs rucksacks) split() groups {
	gs := groups{}
	for i, r := range rs {
		// compute where to put the items
		itemsIndex := i % groupSize
		groupIndex := i / groupSize
		// upsert group
		if groupIndex >= len(gs) {
			gs = append(gs, &group{})
		}
		g := gs[groupIndex]
		// place items into the group
		g[itemsIndex] = r
	}
	return gs
}

type rucksack []item

func (r rucksack) split() (rucksack, rucksack) {
	count := len(r)
	if count%2 != 0 {
		panic("odd number of items")
	}
	left := r[0 : count/2]
	right := r[count/2:]
	return left, right
}

func (is rucksack) index() map[item]bool {
	m := make(map[item]bool)
	for _, i := range is {
		m[i] = true
	}
	return m
}

var empty = item(0)

func (is rucksack) intersect(os rucksack) rucksack {
	intersection := rucksack{}
	idx := is.index()
	for item := range os.index() {
		if _, ok := idx[item]; ok {
			intersection = append(intersection, item)
		}
	}
	return intersection
}

func (rs rucksack) one() (item, error) {
	if len(rs) == 0 {
		return empty, errors.New("none found")
	}
	if len(rs) > 1 {
		return empty, fmt.Errorf("multiple found (%s)", string(rs))
	}
	return rs[0], nil
}
