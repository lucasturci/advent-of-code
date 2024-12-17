package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var vec1, vec2 []int
	var a, b int
	for {
		if n, _ := fmt.Scanf("%d %d", &a, &b); n != 2 {
			break
		}

		vec1 = append(vec1, a)
		vec2 = append(vec2, b)
	}

	sort.Ints(vec1)
	sort.Ints(vec2)

	ans := 0
	for i := 0; i < len(vec1); i++ {
		ans += abs(vec1[i] - vec2[i])
	}
	fmt.Println(ans)
}
