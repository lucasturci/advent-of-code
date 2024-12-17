package main

import (
	"bufio"
	"fmt"
	"os"
)

func assert(condition bool) {
	if !condition {
		panic("assert failed")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	moves := [][]int{
		{1, 0},   // right
		{-1, 0},  // left
		{0, -1},  // up
		{0, 1},   // down
		{1, 1},   // down right
		{-1, 1},  // down left
		{1, -1},  // up right
		{-1, -1}, // up left
	}

	ans := 0
	n := len(grid)
	m := len(grid[0])

	inside := func(i, j, n, m int) bool {
		return i >= 0 && j >= 0 && i < n && j < m
	}

	for i := 0; i < n; i++ {
		assert(len(grid[i]) == m)
		for j := 0; j < m; j++ {
			for _, mov := range moves {
				str := ""
				for k := 0; k < 4; k++ {
					r := i + k*mov[0]
					c := j + k*mov[1]
					if !inside(r, c, n, m) {
						break
					}

					str += string(grid[r][c])
				}
				if str == "XMAS" {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
