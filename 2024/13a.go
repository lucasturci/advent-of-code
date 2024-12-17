package main

import (
	"fmt"
)

const costA = 3
const costB = 1

func solve(a, b, at []int) int {
	men := 100000000
	for i := 0; i <= 100; i++ {
		left := []int{at[0] - a[0]*i, at[1] - a[1]*i}
		if left[0] < 0 || left[1] < 0 {
			continue
		}
		if left[0]%b[0] == 0 && left[1]%b[1] == 0 && left[0]/b[0] == left[1]/b[1] {
			men = min(men, i*costA+(left[0]/b[0])*costB)
		}
	}
	if men == 100000000 {
		return 0
	}
	return men
}

func main() {
	ans := 0
	for {
		a := make([]int, 2)
		b := make([]int, 2)
		at := make([]int, 2)
		var read int
		read, _ = fmt.Scanf("Button A: X+%d, Y+%d\n", &a[0], &a[1])
		if read != 2 {
			break
		}
		read, _ = fmt.Scanf("Button B: X+%d, Y+%d\n", &b[0], &b[1])
		if read != 2 {
			break
		}
		read, _ = fmt.Scanf("Prize: X=%d, Y=%d\n", &at[0], &at[1])
		if read != 2 {
			break
		}
		ans += solve(a, b, at)
		fmt.Scanf("\n")
	}

	fmt.Println(ans)
}
