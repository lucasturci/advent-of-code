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

func calcAns(grid [][]int) int {
	ans := 1
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			offset1 := a * (N/2 + 1)
			offset2 := b * (M/2 + 1)
			sum := 0
			for i := offset1; i < offset1+N/2; i++ {
				for j := offset2; j < offset2+M/2; j++ {
					sum += grid[i][j]
				}
			}
			ans *= sum
		}
	}
	return ans
}

func main() {
	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, M)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		pos, vel := parseLine(line)

		// move thing 100 times forward
		if vel.first() < 0 {
			vel[0] %= M
			vel[0] += M
		}
		if vel.second() < 0 {
			vel[1] %= N
			vel[1] += N
		}

		pos[0] = pos[0] + 100*vel[0]
		pos[0] %= M
		pos[1] = pos[1] + 100*vel[1]
		pos[1] %= N

		grid[pos.second()][pos.first()]++
	}

	ans := calcAns(grid)
	fmt.Println(ans)
}
