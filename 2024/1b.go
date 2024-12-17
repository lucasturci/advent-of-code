package main

import (
	"fmt"
)

func main() {
	var vec []int
	cnt := make(map[int]int)
	var a, b int
	for {
		if n, _ := fmt.Scanf("%d %d", &a, &b); n != 2 {
			break
		}

		vec = append(vec, a)
		cnt[b]++
	}

	ans := 0
	for _, x := range vec {
		ans += x * cnt[x]
	}
	fmt.Println(ans)
}
