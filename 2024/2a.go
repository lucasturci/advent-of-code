package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	var ans int = 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Create a scanner for the line
		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)

		var x, last int
		ok := true
		var lastSign int

		for i := 0; lineScanner.Scan(); i++ {
			fmt.Sscanf(lineScanner.Text(), "%d", &x)

			if i > 1 && sign(x-last) != lastSign {
				ok = false
			}

			if i > 0 {
				if abs(last-x) < 1 || abs(last-x) > 3 {
					ok = false
				}
				lastSign = sign(x - last)
			}

			last = x
		}

		if ok {
			ans++
		}
	}
	fmt.Println(ans)
}
