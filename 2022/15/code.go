package main

import (
	"aoc-in-go/x"
	"log"
	"regexp"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/tidwall/rtree"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	example := len(lines) < 20
	if part2 && !example {
		return nil
	}
	w := world{
		beacons: map[position]*beacon{},
	}
	for _, line := range lines {
		spos, bpos := parseLine(line)
		b := w.addBeacon(bpos)
		w.addSensor(spos, b)
	}

	if part2 {
		maxY := 4000000
		if example {
			maxY = 20
		}
		return w.getDistressBeaconFreq(maxY)
	}

	y := 2000000
	if example {
		y = 10
	}
	return w.getBeaconCoverage(y)
}

type world struct {
	min, max  position
	maxRadius int
	tree      rtree.RTreeG[*sensor]
	sensors   []*sensor
	beacons   map[position]*beacon
}

func (w *world) getDistressBeaconFreq(max int) int {
	possible := []position{}

	// TODO search position in tree,
	// exit early on first found
	// assume only1, return early

	for x := 0; x <= max; x++ {
		for y := 0; y <= max; y++ {
			p := position{x, y}
			if !w.isCovered(p) {
				possible = append(possible, p)
			}
		}
	}
	if len(possible) >= 2 {
		log.Panicf("multiple found: %d", len(possible))
	}
	if len(possible) == 0 {
		panic("none found")
	}
	return possible[0].tuningFreq()
}

func (w *world) getBeaconCoverage(y int) int {
	n := 0
	for x := w.min.x - w.maxRadius; x < w.max.x+w.maxRadius; x++ {
		if w.hasNoBeacon(position{x, y}) {
			n++
		}
	}
	return n
}

func (w *world) hasNoBeacon(p position) bool {
	// check if a beacon is already here
	if _, ok := w.beacons[p]; ok {
		return false
	}
	return w.isCovered(p)
}

func (w *world) isCovered(p position) bool {
	// check if p is inside the radius of any sensor
	for _, s := range w.sensors {
		if s.dist(p) <= s.radius() {
			return true
		}
	}
	return false

}

func (w *world) addBeacon(p position) *beacon {
	if b, ok := w.beacons[p]; ok {
		return b
	}
	w.trackMinMax(p)
	b := &beacon{p}
	w.beacons[p] = b
	return b
}

func (w *world) addSensor(spos position, b *beacon) {
	w.trackMinMax(spos)
	s := &sensor{spos, b}
	w.sensors = append(w.sensors, s)
	if r := s.radius(); r > w.maxRadius {
		w.maxRadius = r
	}
	rect := s.rect()
	w.tree.Insert(rect[0], rect[1], s)
}

func (w *world) trackMinMax(p position) {
	if w.min == (position{}) {
		w.min = p
	}
	if p.x < w.min.x {
		w.min.x = p.x
	}
	if p.y < w.min.y {
		w.min.y = p.y
	}
	if w.max == (position{}) {
		w.max = p
	}
	if p.x > w.max.x {
		w.max.x = p.x
	}
	if p.y > w.max.y {
		w.max.y = p.y
	}
}

type sensor struct {
	position
	closest *beacon
}

func (s *sensor) radius() int {
	return s.position.dist(s.closest.position)
}

func (s *sensor) point() point {
	return point{float64(s.x), float64(s.y)}
}

func (s *sensor) rect() rect {
	centre := s.point()
	r := float64(s.radius())
	min := point{centre[0] - r, centre[1] - r}
	max := point{centre[0] + r, centre[1] + r}
	return rect{min, max}
}

type beacon struct {
	position
}

type position struct {
	x, y int
}

type point [2]float64

type rect [2]point

func (p position) tuningFreq() int {
	return (p.x * 4000000) + p.y
}

// manhatten distance
func (p position) dist(o position) int {
	return x.Abs(p.x-o.x) + x.Abs(p.y-o.y)
}

func parseLine(line string) (sensor, beacon position) {
	const pos = `x=(-?\d+), y=(-?\d+)`
	re := regexp.MustCompile(`Sensor at ` + pos + `: closest beacon is at ` + pos)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 5 {
		panic("invalid line: " + line)
	}
	sensor = position{x.Atoi(matches[1]), x.Atoi(matches[2])}
	beacon = position{x.Atoi(matches[3]), x.Atoi(matches[4])}
	return
}
