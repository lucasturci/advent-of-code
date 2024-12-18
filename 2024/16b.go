package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/lucasturci/everything-go/data-structures/priority_queue"
	"golang.org/x/exp/constraints"
)

type tuple[T constraints.Integer] []T

func (p tuple[T]) sum(pp tuple[T]) tuple[T] {
	return tuple[T]{p[0] + pp[0], p[1] + pp[1], p[2]}
}

func parseInput() [][]byte {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 { // read the moves next
			break
		}
		row := []byte(line)
		grid = append(grid, row)
	}
	return grid
}

func findPosition(grid [][]byte, what byte) tuple[int] {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == what {
				return tuple[int]{i, j}
			}
		}
	}
	panic("Should not arrive here findPosition")
}

func printGrid(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(string(grid[i]))
	}
}

func hash(p tuple[int]) int {
	return 100000*p[0] + 4*p[1] + p[2]
}

type TupleComparator struct{}

func (c TupleComparator) Less(a, b tuple[int]) bool {
	return a[3] < b[3]
}

func dijkstra(grid [][]byte, starts []tuple[int]) map[int]int {
	pq := priority_queue.NewPriorityQueueCustom[tuple[int], TupleComparator]()
	d := make(map[int]int)

	for _, start := range starts {
		pq.Push(append(start, 0))
		d[hash(start)] = 0
	}

	for pq.Size() > 0 {
		u, _ := pq.Top()
		pq.Pop()
		if u[3] > d[hash(u)] {
			continue
		}
		u = u[:3]

		moves := []tuple[int]{
			tuple[int]{0, 1},
			tuple[int]{1, 0},
			tuple[int]{0, -1},
			tuple[int]{-1, 0},
		}
		v := u.sum(moves[u[2]])
		if grid[v[0]][v[1]] != '#' {
			if x, ok := d[hash(v)]; !ok || x > d[hash(u)]+1 {
				d[hash(v)] = d[hash(u)] + 1
				pq.Push(append(v, d[hash(v)]))
			}
		}

		v = slices.Clone(u)
		v[2] = (v[2] + 1) % 4
		if x, ok := d[hash(v)]; !ok || x > d[hash(u)]+1000 {
			d[hash(v)] = d[hash(u)] + 1000
			pq.Push(append(v, d[hash(v)]))
		}

		v = slices.Clone(u)
		v[2] = (v[2] + 3) % 4
		if x, ok := d[hash(v)]; !ok || x > d[hash(u)]+1000 {
			d[hash(v)] = d[hash(u)] + 1000
			pq.Push(append(v, d[hash(v)]))
		}
	}
	return d
}

func main() {
	grid := parseInput()

	start := findPosition(grid, 'S')
	end := findPosition(grid, 'E')

	start = append(start, 0)

	d1 := dijkstra(grid, []tuple[int]{start})
	dist := 10000000
	for k := 0; k < 4; k++ {
		dist = min(dist, d1[hash(append(end, k))])
	}
	d2 := dijkstra(grid, []tuple[int]{append(slices.Clone(end), 0), append(slices.Clone(end), 1), append(slices.Clone(end), 2), append(slices.Clone(end), 3)})
	fmt.Println(dist)

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ok := false
			for k := 0; k < 4; k++ {
				if d1[hash(tuple[int]{i, j, k})]+d2[hash(tuple[int]{i, j, (k + 2) % 4})] == dist {
					ok = true
					grid[i][j] = 'O'
				}
			}
			if ok {
				ans++
			}
		}
	}

	// printGrid(grid)

	fmt.Println(ans)

}
