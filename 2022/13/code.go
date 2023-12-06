package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	lines := strings.Split(input, "\n")
	// parse pairs
	all := items{}
	pairs := pairs{}
	p := &pair{}
	for _, line := range lines {
		if line == "" {
			pairs = append(pairs, p)
			p = &pair{}
			continue
		}
		l := items{}
		l.unmarshal(line)
		if p.left == nil {
			p.left = l
		} else if p.right == nil {
			p.right = l
		} else {
			panic("too many lines")
		}
		all = append(all, l)
	}
	// sort all items with dividiers
	if part2 {
		all.add("[[2]]")
		all.add("[[6]]")
		sort.Sort(all)
		mul := 1
		for i, item := range all {
			if s := item.String(); s == "[[2]]" || s == "[[6]]" {
				mul *= (i + 1)
			}
		}
		return mul
	}
	// sum index of ordered pairs
	sum := 0
	for i, p := range pairs {
		if p.ordered() == right {
			sum += i + 1
		}
	}
	return sum
}

type pairs []*pair

type pair struct {
	left, right items
}

type order int

const (
	wrong   order = -1
	unknown order = 0
	right   order = 1
)

type item interface {
	ordered(item) order
	String() string
}

func (p pair) String() string {
	return fmt.Sprintf("L%s R%s", p.left, p.right)
}

func (p pair) ordered() order {
	return p.left.ordered(p.right)
}

type items []item

func (is items) Len() int {
	return len(is)
}

func (is items) Less(i, j int) bool {
	return is[i].ordered(is[j]) != wrong
}

func (is items) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is items) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	for i, item := range is {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(item.String())
	}
	sb.WriteRune(']')
	return sb.String()
}

func (is *items) unmarshal(line string) {
	var vs []any
	if json.Unmarshal([]byte(line), &vs) != nil {
		panic("invalid line")
	}
	for _, v := range vs {
		is.add(v)
	}
}

func (is *items) add(v any) {
	switch v := v.(type) {
	case string:
		is.unmarshal("[" + v + "]")
	case int:
		*is = append(*is, num(v))
	case float64:
		*is = append(*is, num(v))
	case []any:
		sub := items{}
		for _, v := range v {
			sub.add(v)
		}
		*is = append(*is, sub)
	default:
		log.Panicf("invalid type %T", v)
	}
}

func (is items) ordered(other item) order {
	switch other := other.(type) {
	case items:
		ls := is
		rs := other
		i := 0
		for {
			if i == len(ls) {
				return right
			}
			if i == len(rs) {
				return wrong
			}
			l := ls[i]
			r := rs[i]
			o := l.ordered(r)
			if o != unknown {
				return o
			}
			i++
		}
	case num:
		return is.ordered(items{other})
	}
	panic("unknown item")
}

type num int

func (n num) ordered(other item) order {
	switch other := other.(type) {
	case items:
		return items{n}.ordered(other)
	case num:
		if n < other {
			return right
		}
		if n > other {
			return wrong
		}
		return unknown
	}
	panic("unknown item")
}

func (n num) String() string {
	return fmt.Sprintf("%d", n)
}
