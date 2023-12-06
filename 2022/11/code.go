package main

import (
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
		return "part 2 not solved"
	}
	if true {
		return "part 1 is wrong"
	}
	// init game
	g := game{
		round:   0,
		monkeys: []*monkey{},
	}
	// parse monkeys
	for input != "" {
		m := &monkey{}
		input = m.unmarshal(input)
		g.monkeys = append(g.monkeys, m)
	}
	// play N rounds
	rounds := 20
	if part2 {
		rounds = 10_000
	}
	// simulate
	const debug = true
	for n := 1; n <= rounds; n++ {
		g.playRound(part2)
		if n%100 == 0 {
			fmt.Printf("Played round %d\n", g.round)
			if debug {
				g.showItems()
			}
		}
	}
	// return result
	fmt.Printf("Completed %d rounds\n", g.round)
	return g.monkeyBusiness()
}

type game struct {
	round   int
	monkeys []*monkey
}

func (g *game) playRound(div3 bool) {
	for _, m := range g.monkeys {
		items := m.items
		m.items = nil
		m.inspections += len(items)
		for _, v := range items {
			m.op.applyTo(v)
			if div3 {
				v.Div(v, big.NewInt(3))
			}
			mod := big.Int{}
			mod.Mod(v, m.test.div)
			divs := mod.Int64() == 0
			var id int
			if divs {
				id = m.test.trueId
			} else {
				id = m.test.falseId
			}
			next := g.monkeys[id]
			next.items = append(next.items, v)
		}
	}
	g.round++
}

func (g *game) showItems() {
	fmt.Printf("Round %d\n", g.round)
	for _, m := range g.monkeys {
		fmt.Printf("Monkey %d: %v (inspections %d)\n", m.id, m.items, m.inspections)
	}
}

func (g *game) monkeyBusiness() int {
	by := byInspections(g.monkeys)
	sort.Sort(by)
	return by[0].inspections * by[1].inspections
}

type monkey struct {
	id          int
	inspections int
	items       []*big.Int
	op          op
	test        test
}

type op struct {
	add, mul, sq bool
	value        *big.Int
}

func (o op) applyTo(value *big.Int) {
	if o.sq {
		value.Mul(value, value)
	} else if o.add {
		value.Add(value, o.value)
	} else if o.mul {
		value.Mul(value, o.value)
	} else {
		panic("invalid op")
	}
}

type test struct {
	div     *big.Int
	trueId  int
	falseId int
}

var monkeyRe = regexp.MustCompile(`(?m)Monkey (\d+):\n` +
	`  Starting items: (\d+(, \d+)*)\n` +
	`  Operation: new = old ([\+\*]) (\d+|old)\n` +
	`  Test: divisible by (\d+)\n` +
	`    If true: throw to monkey (\d+)\n` +
	`    If false: throw to monkey (\d+)\n?\n?`)

func (m *monkey) unmarshal(input string) string {
	loc := monkeyRe.FindStringIndex(input)
	if len(loc) == 0 {
		if len(input) > 0 {
			panic("expected end of string")
		}
		return ""
	}
	if loc[0] != 0 {
		panic("expected start of string")
	}
	sub := input[loc[0]:loc[1]]
	rest := input[loc[1]:]
	values := monkeyRe.FindStringSubmatch(sub)
	m.id = atoi(values[1])
	m.inspections = 0
	for _, v := range strings.Split(values[2], ", ") {
		m.items = append(m.items, atobi(v))
	}
	if values[5] == "old" {
		m.op.sq = true
	} else {
		m.op.add = values[4] == "+"
		m.op.mul = values[4] == "*"
		m.op.value = atobi(values[5])
	}
	m.test.div = atobi(values[6])
	m.test.trueId = atoi(values[7])
	m.test.falseId = atoi(values[8])
	return rest
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func atobi(s string) *big.Int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return big.NewInt(n)
}

type byInspections []*monkey

func (m byInspections) Len() int {
	return len(m)
}

func (m byInspections) Less(i, j int) bool {
	return m[i].inspections > m[j].inspections
}

func (m byInspections) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
