package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	n := len(str)

	tot := 0
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(string(str[i]))
		tot += x
	}

	vec := make([]int, tot+1)
	cur := 0
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(string(str[i]))
		for j := 0; j < x; j++ {
			if i%2 == 0 {
				vec[cur] = i / 2
			} else {
				vec[cur] = -1
			}
			cur++
		}
	}

	for l, r := 0, tot-1; l < r; r-- {
		if vec[r] >= 0 {
			for l < r && vec[l] >= 0 {
				l++
			}
			vec[l] = vec[r]
			if l < r {
				vec[r] = -1
			}
		}
	}

	var ans int64
	for i := 0; i < tot; i++ {
		if vec[i] == -1 {
			continue
		}
		ans += int64(i) * int64(vec[i])
	}
	fmt.Println(ans)
}
