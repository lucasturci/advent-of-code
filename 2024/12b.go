package main

import (
	"bufio"
	"fmt"
	"os"
)

var moves = [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

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

	inside := func(i, j int) bool {
		if i < 0 || i >= n || j < 0 || j >= m {
			return false
		}
		return true
	}

	var flood_fill func(vis map[int64]bool, i, j int, col byte)
	flood_fill = func(vis map[int64]bool, i, j int, col byte) {
		if !inside(i, j) || grid[i][j] != col {
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

	// calculate vertices
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k, _ := range moves {
				p, q := i+moves[k][0], j+moves[k][1]
				r, s := i+moves[(k+1)%4][0], j+moves[(k+1)%4][1]
				dif := func(p, q int) bool { // different than i, j
					return !inside(p, q) || grid[p][q] != grid[i][j]
				}
				if dif(p, q) && dif(r, s) { // convex vertex
					ans += area[i][j]
					if inside(p+r-i, q+s-j) && grid[p][q] == grid[r][s] { // calculate its concave pair
						p, q := p+r-i, q+s-j
						if grid[p][q] == grid[r][s] {
							ans += area[p][q]
						}
					}
				}
			}
		}
	}

	fmt.Println(ans)

}
