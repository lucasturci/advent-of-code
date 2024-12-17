package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseStr(str string) int {
	re := regexp.MustCompile(`^mul\((\d{1,3}),(\d{1,3})\)`)

	ans := 0
	res := re.FindAllStringSubmatch(str, -1)
	for _, match := range res {
		if len(match) != 3 {
			fmt.Errorf("Something didn't work")
		}
		x, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Errorf("Conversion did not work")
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Errorf("Conversion did not work")
		}
		ans += x * y
	}
	return ans
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	ans := 0
	on := true
	for scanner.Scan() {
		line := scanner.Text()

		n := len(line)
		for i := 0; i < n; i++ {
			if line[i:min(n, i+4)] == "do()" {
				on = true
			}
			if line[i:min(n, i+7)] == "don't()" {
				on = false
			}

			snippet := line[i:min(n, i+12)]
			if on {
				ans += parseStr(snippet)
			}
		}

	}
	fmt.Println(ans)
}
