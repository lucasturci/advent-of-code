package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int

// didn't make dp, apparently the input didn't require it
func solve(grid []string, i, j int, x byte) int64 {
	if i < 0 || i >= n || j < 0 || j >= m {
		return 0
	}
	if grid[i][j] != x {
		return 0
	}

	if grid[i][j] == '9' {
		return 1
	}
	moves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var ans int64 = 0
	for _, move := range moves {
		ans += solve(grid, i+move[0], j+move[1], x+1)
	}
	return ans
}

func main() {
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	n = len(grid)
	m = len(grid[0])

	var ans int64 = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				ans += solve(grid, i, j, '0')
			}
		}
	}

	fmt.Println(ans)
}
