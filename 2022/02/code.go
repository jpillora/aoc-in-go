package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		pair := strings.SplitN(line, " ", 2)
		var op move
		op.unmarshal('A', pair[0])
		if op.String() == "unknown" {
			panic("invalid opponent move")
		}
		if part2 {
			var o outcome
			o.unmarshal(pair[1])
			me := op.want(o)
			sum += o.score() + me.score()
		} else {
			var me move
			me.unmarshal('X', pair[1])
			s := op.play(me).score()
			sum += me.score() + s
		}
	}
	return sum
}

type outcome byte

const (
	lose outcome = iota
	draw
	win
)

func (o *outcome) unmarshal(s string) {
	*o = outcome(delta('X', s))
}

func (o outcome) String() string {
	switch o {
	case lose:
		return "lose"
	case draw:
		return "draw"
	case win:
		return "win"
	}
	return "unknown"
}

func (o outcome) score() int {
	switch o {
	case lose:
		return 0
	case draw:
		return 3
	case win:
		return 6
	}
	panic("invalid outcome")
}

type move byte

const (
	rock move = iota
	paper
	scissors
)

func (m *move) unmarshal(base byte, s string) {
	*m = move(delta(base, s))
}

func (m move) String() string {
	switch m {
	case rock:
		return "rock"
	case paper:
		return "paper"
	case scissors:
		return "scissors"
	}
	return "unknown"
}

func (m move) score() int {
	return int(m) + 1
}

func (m move) play(against move) outcome {
	if m == against {
		return draw
	}
	if byte(m) == up(byte(against), 3) {
		return lose
	}
	if byte(m) == down(byte(against), 3) {
		return win
	}
	panic(fmt.Sprintf("invalid moves: %d plays %d", m, against))
}

func (m move) want(outcome outcome) move {
	switch outcome {
	case lose:
		return move(down(byte(m), 3))
	case draw:
		return m
	case win:
		return move(up(byte(m), 3))
	}
	panic("invalid outcome")
}

func delta(base byte, m string) byte {
	if len(m) != 1 {
		panic("invalid delta")
	}
	return m[0] - base
}

func up(n, mod byte) byte {
	return (n + 1) % mod
}

func down(n, mod byte) byte {
	if n == 0 {
		return mod - 1
	}
	return (n - 1) % mod
}
