package main

import (
	"container/heap"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	top := 1
	if part2 {
		top = 3
	}
	sum := 0
	elves := &minHeap{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			heap.Push(elves, sum)
			sum = 0
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += n
	}
	elfSum := 0
	for n := 1; n <= top; n++ {
		elfSum += heap.Pop(elves).(int)
	}
	return elfSum
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
