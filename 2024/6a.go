package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]byte{}

	var i, j int
	var n, m int
	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))

		for k, c := range line {
			if c == '^' {
				i, j = cnt, k
			}
		}
		m = len(grid)
		cnt++
	}
	n = cnt

	ans := 0
	dir := 0
	moves := [][]int{
		[]int{-1, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
	}

	inside := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < n && j < m
	}

	for {
		grid[i][j] = 'X'
		r, c := i+moves[dir][0], j+moves[dir][1]
		if inside(r, c) && grid[r][c] == '#' {
			dir = (dir + 1) % 4
			continue
		} else if !inside(r, c) {
			break
		}
		i, j = r, c
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'X' {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
