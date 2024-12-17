package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type toposort struct {
	g    [][]int
	topo map[int]int
	cnt  int
}

func NewToposort(n int) toposort {
	return toposort{
		g:    make([][]int, n+1),
		topo: make(map[int]int, n+1),
	}
}

func (t *toposort) addEdge(u, v int) {
	t.g[v] = append(t.g[v], u)
}

func (t *toposort) run() map[int]int {
	for i := range t.g {
		t.dfs(i)
	}

	return t.topo
}

func (t *toposort) dfs(u int) {
	if t.topo[u] > 0 {
		return
	}
	if t.topo[u] == -1 { // cycle
		return
	}
	t.topo[u] = -1
	for _, v := range t.g[u] {
		t.dfs(v)
	}
	t.cnt++
	t.topo[u] = t.cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var ans int = 0
	var mai int = 0
	var topo map[int]int
	var edges [][]int
	for scanner.Scan() {
		line := scanner.Text()

		var x, y int
		if strings.Contains(line, "|") {
			fmt.Sscanf(line, "%d|%d", &x, &y)
			edges = append(edges, []int{x, y})
			mai = max(mai, max(x, y))
		} else if strings.Contains(line, ",") {
			words := strings.Split(line, ",")

			vec := []int{}
			for _, w := range words {
				x, _ := strconv.Atoi(w)
				vec = append(vec, x)
			}

			ok := true
			for i := 1; i < len(vec); i++ {
				if topo[vec[i]] < topo[vec[i-1]] {
					fmt.Println(vec[i], topo[vec[i]])
					fmt.Println(vec[i-1], topo[vec[i-1]])
					fmt.Println()
					ok = false
					break
				}
			}
			if ok {
				ans += vec[len(vec)/2]
			}

			// quadratic approach first
		} else {
			// run toposort
			toposort := NewToposort(mai)
			for _, ed := range edges {
				toposort.addEdge(ed[0], ed[1])
			}
			topo = toposort.run()

			fmt.Println(topo)
		}
	}
	fmt.Println(ans)
}
