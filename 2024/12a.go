package main

import (
	"bufio"
	"fmt"
	"os"
)

var moves = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
	grid := []string{}
	area := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, line)
		area = append(area, make([]int, len(line)))
	}

	n, m := len(grid), len(grid[0])

	var flood_fill func(vis map[int64]bool, i, j int, col byte)
	flood_fill = func(vis map[int64]bool, i, j int, col byte) {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] != col {
			return
		}
		_, ok := vis[(int64(i)<<32)|int64(j)]
		if ok {
			return
		}
		vis[(int64(i)<<32)|int64(j)] = true
		for _, move := range moves {
			p, q := i+move[0], j+move[1]
			flood_fill(vis, p, q, col)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			vis := make(map[int64]bool)
			flood_fill(vis, i, j, grid[i][j])
			for k, _ := range vis {
				area[k>>32][k&((1<<32)-1)] = len(vis)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, move := range moves {
				p, q := i+move[0], j+move[1]
				if p < 0 || q < 0 || p >= n || q >= m || grid[p][q] != grid[i][j] {
					ans += area[i][j]
				}
			}
		}
	}

	fmt.Println(ans)

}
