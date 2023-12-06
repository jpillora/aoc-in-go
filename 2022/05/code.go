package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	stacks := stacks{map[int]*stack{}}
	build := true
	actions := []action{}
	for i, line := range strings.Split(input, "\n") {
		if i == 0 && line == "" {
			continue
		}
		if line == "" {
			build = false
			continue
		}
		if build {
			stacks.decode(line)
		} else {
			a := action{}
			a.unmarshal(line)
			a.ordered = part2
			actions = append(actions, a)
		}
	}
	for _, a := range actions {
		stacks.perform(a)
	}
	return string(stacks.top())
}

type stacks struct {
	stacks map[int]*stack
}

func (ss stacks) String() string {
	p := []string{}
	for i := 1; ss.stacks[i] != nil; i++ {
		p = append(p, fmt.Sprintf("%d:%s", i, ss.stacks[i].String()))
	}
	return fmt.Sprintf("stacks{%s}", strings.Join(p, ", "))
}

func (ss stacks) decode(line string) {
	// every 4 chars, we check for a crate
	for i := 1; i < len(line); i += 4 {
		isCreate := line[i-1] == '[' && line[i+1] == ']'
		c := crate(line[i])
		if !isCreate {
			continue
		}
		// we choose a stack id based on the index
		sid := (i / 4) + 1
		s, ok := ss.stacks[sid]
		if !ok {
			s = &stack{}
			ss.stacks[sid] = s
		}
		// built top down, so it is reversed
		s.stack = append(s.stack, c)
	}
}

func (ss stacks) perform(a action) {
	from := ss.stacks[a.from]
	to := ss.stacks[a.to]
	if a.ordered {
		from.move(a.count, to)
	} else {
		from.singleMove(a.count, to)
	}
}

func (ss stacks) top() []crate {
	crates := []crate{}
	for i := 1; ss.stacks[i] != nil; i++ {
		crates = append(crates, ss.stacks[i].stack[0])
	}
	return crates
}

type crate = rune

type stack struct {
	stack []crate
}

func (s stack) String() string {
	return "[" + string(s.stack) + "]"
}

func (s *stack) singleMove(count int, to *stack) {
	from := s
	for n := 1; n <= count; n++ {
		from.move(1, to)
	}
}

func (s *stack) move(count int, to *stack) {
	from := s
	tmp := make([]rune, count)
	copy(tmp, from.stack[0:count])
	from.stack = from.stack[count:]
	to.stack = append(tmp, to.stack...)
}

type action struct {
	ordered  bool
	count    int
	from, to int
}

func (a action) String() string {
	return fmt.Sprintf("action{move %d from %d to %d}", a.count, a.from, a.to)
}

var re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func (a *action) unmarshal(s string) {
	m := re.FindStringSubmatch(s)
	if len(m) != 4 {
		log.Panicf("invalid action: %s", s)
	}
	a.count = atoi(m[1])
	a.from = atoi(m[2])
	a.to = atoi(m[3])
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
