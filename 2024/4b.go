package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	moves := [][]int{
		{1, 1},  // down right
		{-1, 1}, // down left
	}

	ans := 0
	n := len(grid)
	m := len(grid[0])

	for i := 1; i+1 < n; i++ {
		for j := 1; j+1 < m; j++ {
			if grid[i][j] == 'A' {
				words := make([]string, 2)
				for k := range moves {
					for r, c, l := i-moves[k][0], j-moves[k][1], 0; l < 3; l++ {
						words[k] += string(grid[r][c])

						r += moves[k][0]
						c += moves[k][1]
					}
				}
				if (words[0] == "MAS" || words[0] == "SAM") && (words[1] == "MAS" || words[1] == "SAM") {
					ans++
				}

			}
		}
	}
	fmt.Println(ans)
}
