package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

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
			if grid[i][j] == '.' {
				continue
			}
			for ii := 0; ii < n; ii++ {
				for jj := 0; jj < m; jj++ {
					if (ii == i && jj == j) || grid[ii][jj] != grid[i][j] {
						continue
					}

					vec := []int{ii - i, jj - j}
					g := gcd(abs(ii-i), abs(jj-j))
					vec[0] /= g
					vec[1] /= g
					for iii, jjj := i, j; inside(iii, jjj); iii, jjj = iii+vec[0], jjj+vec[1] {
						vis[iii][jjj] = true
					}
					for iii, jjj := i, j; inside(iii, jjj); iii, jjj = iii-vec[0], jjj-vec[1] {
						vis[iii][jjj] = true
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
