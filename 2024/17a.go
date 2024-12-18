package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() (int, int, int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	var A, B, C int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		var key byte
		var val int
		fmt.Sscanf(line, "Register %c: %d", &key, &val)
		if key == 'A' {
			A = val
		} else if key == 'B' {
			B = val
		} else {
			C = val
		}
	}

	scanner.Scan()
	line := scanner.Text()
	line = strings.Replace(line, "Program: ", "", -1)
	sl := strings.Split(line, ",")
	program := []int{}
	for _, c := range sl {
		x, _ := strconv.Atoi(c)
		program = append(program, x)
	}
	return A, B, C, program
}

func main() {
	A, B, C, program := parseInput()

	combo := func(x int) int {
		switch {
		case x <= 3:
			return x
		case x == 4:
			return A
		case x == 5:
			return B
		case x == 6:
			return C
		default:
			panic("Combo 7 should not appear")
		}
	}

	ip := 0 // instructino ptr
	res := []int{}
	do := map[int]func(int){
		0: func(op int) {
			A = A >> combo(op)
		},
		1: func(op int) {
			B = B ^ op
		},
		2: func(op int) {
			B = combo(op) & 7
		},
		3: func(op int) {
			if A == 0 {
				return
			}
			ip = op
		},
		4: func(_ int) {
			B = B ^ C
		},
		5: func(op int) {
			res = append(res, combo(op)&7)
		},
		6: func(op int) {
			B = A >> combo(op)
		},
		7: func(op int) {
			C = A >> combo(op)
		},
	}

	for {
		if ip >= len(program) {
			break // halt
		}
		opcode, op := program[ip], program[ip+1]
		do[opcode](op)
		if opcode != 3 || A == 0 { // jnz
			ip += 2
		}
	}

	for i, x := range res {
		fmt.Printf("%v", x)
		if i+1 != len(res) {
			fmt.Printf(",")
		}
	}
	fmt.Println()

}
