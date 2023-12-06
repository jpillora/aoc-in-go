package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/ansi"
	"github.com/jpillora/puzzler/harness/aoc"
)

const draw = false

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	// parse input
	w := world{}
	for _, line := range strings.Split(input, "\n") {
		w = append(w, tiles(line))
	}
	// dykstra's algorithm
	g := newGraph(w)
	// choose node to begin at
	begin := g.E
	begin.cost = 0
	batch := []*node{begin}
	count := 1
	for len(batch) > 0 {
		next := []*node{}
		for _, node := range batch {
			if node.visited {
				continue
			}
			for _, adj := range g.adjacent(node) {
				// update cost (maybe shorter path to visted)
				adj.cost = node.cost + 1
				adj.parent = node
				// add unvisted to next batch
				next = append(next, adj)
			}
			node.visited = true
		}
		batch = next
		count++
	}
	// part 1 target is S
	target := g.S
	// part 2 target is min 'a'
	if part2 {
		for _, n := range g.nodes {
			if n.tile == 'a' && n.cost != -1 && n.cost < target.cost {
				target = n
			}
		}
	}
	// draw path to target from E
	if draw {
		g.printPath(target)
	}
	return target.cost
}

type tile rune

func (t tile) String() string {
	return string(t)
}

func (t tile) isStart() bool {
	return t == 'S'
}

func (t tile) isEnd() bool {
	return t == 'E'
}

func (t tile) elevation() int {
	e := t
	if t.isStart() {
		e = 'a'
	} else if t.isEnd() {
		e = 'z'
	}
	return int(e - 'a' + 1)
}

func (t tile) canReach(t2 tile) bool {
	return t.elevation() >= t2.elevation()-1
}

type tiles []tile

type world []tiles

func (w world) adjacent(from position) []position {
	ns := []position{}
	for _, to := range []position{
		{from.i - 1, from.j},
		{from.i, from.j - 1},
		{from.i, from.j + 1},
		{from.i + 1, from.j},
	} {
		if to.i < 0 || to.i >= len(w) {
			continue
		}
		if to.j < 0 || to.j >= len(w[to.i]) {
			continue
		}
		ns = append(ns, to)
	}
	return ns
}

type graph struct {
	world world
	S, E  *node
	nodes map[position]*node
}

func newGraph(w world) *graph {
	g := &graph{
		world: w,
		nodes: map[position]*node{},
	}
	for i, row := range w {
		for j, t := range row {
			n := newNode(t, position{i, j})
			if t.isStart() {
				g.S = n
			} else if t.isEnd() {
				g.E = n
			}
			g.add(n)
		}
	}
	if g.S == nil {
		panic("no start tile")
	}
	if g.E == nil {
		panic("no end tile")
	}
	return g
}

func (g *graph) add(n *node) {
	g.nodes[n.position] = n
}

func (g *graph) adjacent(n *node) []*node {
	ns := []*node{}
	for _, p := range g.world.adjacent(n.position) {
		t := g.world[p.i][p.j]
		if t.canReach(n.tile) {
			n := g.nodes[position{p.i, p.j}]
			if !n.visited {
				ns = append(ns, n)
			}
		}
	}
	return ns
}

func (g *graph) printPath(to *node) {
	// capture path
	path := map[position]bool{}
	n := to
	for n != nil && !path[n.position] {
		path[n.position] = true
		n = n.parent
	}
	// draw world
	for i, row := range g.world {
		for j, t := range row {
			p := position{i, j}
			s := string(t)
			if path[p] {
				s = ansi.Green.String(s)
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}

type position struct {
	i, j int
}

func (p position) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

type node struct {
	tile
	position
	visited bool
	cost    int
	parent  *node
}

func newNode(t tile, p position) *node {
	return &node{
		tile:     t,
		position: p,
		cost:     -1,
	}
}
