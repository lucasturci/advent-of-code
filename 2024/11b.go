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

func hh(x, steps int) int64 {
	return (int64(x) << 32) | int64(steps)
}

func solve2(dp map[int64]int64, x, steps int) int64 {
	if x == -1 {
		return 0
	}
	if steps == 0 {
		return 1
	}
	if dp[hh(x, steps)] == 0 {
		a, b := transform(x)
		dp[hh(x, steps)] = solve2(dp, a, steps-1) + solve2(dp, b, steps-1)
	}
	return dp[hh(x, steps)]
}

func main() {
	vec := []int{3, 386358, 86195, 85, 1267, 3752457, 0, 741}

	var ans int64 = 0
	dp := make(map[int64]int64)
	for _, x := range vec {
		ans += solve2(dp, x, 75)
	}

	fmt.Println(ans)
}
