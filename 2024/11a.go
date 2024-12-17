package main

import (
	"fmt"
	"strconv"
)

func numberLen(x int) int {
	tot := 0
	for x > 0 {
		tot++
		x /= 10
	}
	return tot
}

func transform(x int) (int, int) {
	if x == 0 {
		return 1, -1
	}
	if numberLen(x)%2 == 0 {
		str := fmt.Sprintf("%d", x)
		s := str[:len(str)/2]
		t := str[len(str)/2:]
		a, _ := strconv.Atoi(s)
		b, _ := strconv.Atoi(t)
		return a, b
	}
	return x * 2024, -1
}

func main() {
	vec := []int{3, 386358, 86195, 85, 1267, 3752457, 0, 741}

	for i := 0; i < 25; i++ {
		nxt := []int{}
		for _, x := range vec {
			a, b := transform(x)
			nxt = append(nxt, a)
			if b != -1 {
				nxt = append(nxt, b)
			}
		}
		vec = nxt
	}

	fmt.Println(len(vec))
}
