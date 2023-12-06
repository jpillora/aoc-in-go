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
	if part2 {
		return nil
	}
	vm := machine{
		registerX: 1,
	}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if line == "noop" {
			vm.tick()
		} else if strings.HasPrefix(line, "addx ") {
			num := strings.TrimPrefix(line, "addx ")
			n, _ := strconv.Atoi(num)
			vm.tick()
			vm.tick()
			vm.registerX += n
		} else {
			panic("invalid instruction: " + line)
		}
	}
	vm.renderScreen()
	return vm.signalStrength
}

const vwidth = 40
const vheight = 6

type machine struct {
	cycle          int
	registerX      int
	signalStrength int
	vbuffer        [vheight][vwidth]bool
}

func (m *machine) tick() {
	// start
	cycle := m.cycle
	// CRT tick
	vrow := cycle / vwidth
	vcol := m.cycle % vwidth
	scol := m.registerX
	lit := scol-1 <= vcol && vcol <= scol+1
	m.vbuffer[vrow][vcol] = lit
	// signal strength tick
	if (cycle+20)%vwidth == 0 {
		m.signalStrength += cycle * m.registerX
	}
	// fin
	m.cycle = cycle + 1
}

func (m *machine) renderScreen() {
	println()
	for row := 0; row < vheight; row++ {
		for col := 0; col < vwidth; col++ {
			if m.vbuffer[row][col] {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}
}
