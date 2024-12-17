package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 103
const M = 101

type Pair[T any] []T

func NewPair[T any](a, b T) Pair[T] {
	return []T{a, b}
}

func (p Pair[T]) first() T {
	return p[0]
}

func (p Pair[T]) second() T {
	return p[1]
}

func parseLine(line string) (Pair[int], Pair[int]) {
	var a, b, c, d int
	fmt.Sscanf(line, "p=%d,%d v=%d,%d", &a, &b, &c, &d)
	pos := NewPair(a, b)
	vel := NewPair(c, d)
	return pos, vel
}

func printTree(grid [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if grid[i][j] != 0 {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	grid := make([][]int, N)
	positions, velocities := make([]Pair[int], 0), make([]Pair[int], 0)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, M)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		pos, vel := parseLine(line)

		if vel.first() < 0 {
			vel[0] %= M
			vel[0] += M
		}
		if vel.second() < 0 {
			vel[1] %= N
			vel[1] += N
		}
		positions = append(positions, pos)
		velocities = append(velocities, vel)
		grid[pos[1]][pos[0]]++
	}

	// move thing one at a time and print when we see a pattern
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(positions); j++ {
			pos := positions[j]
			vel := velocities[j]

			grid[pos[1]][pos[0]]--

			pos[0] = pos[0] + vel[0]
			pos[0] %= M
			pos[1] = pos[1] + vel[1]
			pos[1] %= N
			grid[pos[1]][pos[0]]++
		}

		mai := 0
		for j := 0; j < N; j++ {
			// longest line of XXXXXX
			tot := 0
			for k := 0; k < M; k++ {
				if grid[j][k] > 0 {
					tot++
				} else {
					tot = 0
				}
				mai = max(mai, tot)
			}
		}
		if mai > 10 {
			printTree(grid)
			fmt.Println(i + 1)
		}

	}

}
