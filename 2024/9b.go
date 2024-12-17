package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/lucasturci/everything-go/data-structures/heap"
)

func ptr[T any](value T) *T {
	return &value
}

func assert(cond bool) {
	if !cond {
		fmt.Fprintln(os.Stderr, "Failed assert")
		os.Exit(1)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	n := len(str)

	free := make(map[int]*heap.MinHeap[int])
	for i := 0; i < 10; i++ {
		free[i] = ptr(heap.NewMinHeap[int]())
	}

	tot := 0
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(string(str[i]))
		tot += x
	}

	vec := make([]int, tot)
	pos := make([]int, len(str))
	cur := 0
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(string(str[i]))
		if i%2 == 1 {
			free[x].Push(cur)
		}
		pos[i] = cur
		for j := 0; j < x; j++ {
			if i%2 == 0 {
				vec[cur] = i / 2
			} else {
				vec[cur] = -1
			}
			cur++
		}
	}

	for i := n - 1; i >= 0; i-- {
		if i%2 == 1 {
			continue
		}
		x, _ := strconv.Atoi(string(str[i]))
		assert(x != 0)
		men, y := pos[i], -1
		for j := x; j < 10; j++ {
			if free[j].Size() == 0 {
				continue
			}
			if k, _ := free[j].Top(); k < pos[i] {
				if k < men {
					men = k
					y = j
				}
			}
		}
		if y >= 0 {
			free[y].Pop()
			for p := men; p < men+x; p++ { // paint new thing
				assert(vec[p] == -1)
				vec[p] = i / 2
			}

			for p := pos[i]; p < pos[i]+x; p++ { // unpaint old thing
				assert(vec[p] == i/2)
				vec[p] = -1
			}

			// update min heap
			free[y-x].Push(men + x)
		}
	}

	var ans int64
	for i := 0; i < tot; i++ {
		if vec[i] == -1 {
			continue
		}
		ans += int64(i) * int64(vec[i])
	}
	// fmt.Println(vec)
	fmt.Println(ans)
}
