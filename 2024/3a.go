package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		res := re.FindAllStringSubmatch(line, -1)
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
	}
	fmt.Println(ans)
}
