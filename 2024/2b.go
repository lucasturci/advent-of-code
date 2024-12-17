package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sign(n int) int {
	if n == 0 {
		return 1
	}
	return n / abs(n)
}

func isValidSequence(numbers []int) bool {
	if len(numbers) <= 1 {
		return true
	}

	lastSign := 0
	for i := 1; i < len(numbers); i++ {
		if i > 1 && sign(numbers[i]-numbers[i-1]) != lastSign {
			return false
		}

		if abs(numbers[i]-numbers[i-1]) < 1 || abs(numbers[i]-numbers[i-1]) > 3 {
			return false
		}
		lastSign = sign(numbers[i] - numbers[i-1])
	}
	return true
}

func main() {
	var ans int = 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Create a scanner for the line
		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)

		var x int

		vec := make([]int, 0)

		for i := 0; lineScanner.Scan(); i++ {
			fmt.Sscanf(lineScanner.Text(), "%d", &x)

			vec = append(vec, x)
		}

		ok := false
		for i := 0; i < len(vec); i++ {
			vec2 := slices.Clone[[]int](vec[:i])
			vec2 = append(vec2, vec[i+1:]...)
			if isValidSequence(vec2) {
				ok = true
				break
			}
		}
		if ok {
			ans++
		}
	}
	fmt.Println(ans)
}
