package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func assert(cond bool) {
	if !cond {
		fmt.Errorf("Assert failed")
		os.Exit(1)
	}
}

func test(vec []int, testValue int, ch chan int64) {
	n := len(vec)
	for i := 0; i < int(math.Pow(3, float64(n-1))); i++ {
		val := int64(vec[0])
		cur := i
		for k := 1; k < n; k++ {
			if cur%3 == 1 {
				val += int64(vec[k])
			} else if cur%3 == 2 {
				val *= int64(vec[k])
			} else {
				val *= int64(math.Pow10(len(fmt.Sprintf("%d", vec[k]))))
				val += int64(vec[k])
			}
			cur /= 3
		}
		if val == int64(testValue) {
			ch <- int64(testValue)
			return
		}
	}
	ch <- 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ans := int64(0)

	ch := make(chan int64)

	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()

		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)

		lineScanner.Scan()
		testValueStr := lineScanner.Text()
		assert(testValueStr[len(testValueStr)-1:] == ":")

		testValue, _ := strconv.Atoi(testValueStr[:len(testValueStr)-1])

		vec := []int{}
		for lineScanner.Scan() {
			word := lineScanner.Text()
			x, _ := strconv.Atoi(word)
			vec = append(vec, x)
		}

		// brute force

		if len(vec) == 1 && vec[0] == testValue {
			ans += int64(testValue)
		} else {
			cnt++
			go test(vec, testValue, ch)
		}
	}

	for i := 0; i < cnt; i++ {
		x := <-ch
		ans += x
	}
	fmt.Println(ans)
}
