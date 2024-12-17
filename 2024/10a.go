package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lucasturci/everything-go/data-structures/bitset"
)

var n, m int
var dp []bitset.Bitset
var vis []bool

func hash(i, j int, x byte) int {
	return 10*(i*n+j) + int(x-'0')
}

func solve(grid []string, i, j int, x byte) bitset.Bitset {
	if i < 0 || i >= n || j < 0 || j >= m {
		return bitset.New(n * m)
	}
	if grid[i][j] != x {
		return bitset.New(n * m)
	}

	if vis[hash(i, j, x)] == true {
		return dp[hash(i, j, x)]
	}

	if grid[i][j] == '9' {
		b := bitset.New(n * m)
		b.Set(i*m + j)
		return b
	}
	moves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	ans := bitset.New(n * m)
	for _, move := range moves {
		ans = bitset.Union(solve(grid, i+move[0], j+move[1], x+1), ans)
	}
	dp[hash(i, j, x)] = ans
	vis[hash(i, j, x)] = true

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

	dp = make([]bitset.Bitset, n*m*10)
	vis = make([]bool, n*m*10)

	var ans int = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				set := solve(grid, i, j, '0')
				ans += set.Count()
			}
		}
	}

	fmt.Println(ans)
}
