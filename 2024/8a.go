package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var grid [][]byte
	var vis [][]bool

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
		vis = append(vis, make([]bool, len(line)))
	}

	n := len(grid)
	m := len(grid[0])

	inside := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < n && j < m
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for ii := 0; ii < n; ii++ {
				for jj := 0; jj < m; jj++ {
					if ii == i && jj == j {
						continue
					}

					vec1 := []int{ii - i, jj - j}
					vec2 := []int{vec1[0] * 2, vec1[1] * 2}

					pos1 := []int{i + vec1[0], j + vec1[1]}
					pos2 := []int{i + vec2[0], j + vec2[1]}
					if inside(pos2[0], pos2[1]) && grid[pos1[0]][pos1[1]] != '.' && grid[pos1[0]][pos1[1]] == grid[pos2[0]][pos2[1]] {
						vis[i][j] = true
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
