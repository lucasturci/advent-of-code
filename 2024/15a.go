package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/constraints"
)

type pair[T constraints.Integer] []T

func (p pair[T]) sum(pp pair[T]) pair[T] {
	return pair[T]{p[0] + pp[0], p[1] + pp[1]}
}

func parseInput() ([][]byte, string) {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 { // read the moves next
			break
		}
		row := []byte(line)
		grid = append(grid, row)
	}

	var moves string
	for scanner.Scan() {
		line := scanner.Text()
		moves += line
	}
	return grid, moves
}

func findPosition(grid [][]byte) pair[int] {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' {
				return pair[int]{i, j}
			}
		}
	}
	panic("Should not arrive here findPosition")
}

// tries to move and returns if position is available or not
func shift(grid [][]byte, pos pair[int], delta pair[int]) bool {
	newPos := pos.sum(delta)

	if grid[pos[0]][pos[1]] == '#' {
		return false
	}

	if grid[pos[0]][pos[1]] == '.' {
		return true
	}

	ret := shift(grid, newPos, delta)
	if ret {
		grid[newPos[0]][newPos[1]] = grid[pos[0]][pos[1]]
		grid[pos[0]][pos[1]] = '.'
	}

	return ret
}

func calcAns(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'O' {
				ans += 100*i + j
			}
		}
	}
	return ans
}

func printGrid(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(string(grid[i]))
	}
}

func main() {
	grid, moves := parseInput()

	pos := findPosition(grid)

	charToMove := map[rune]pair[int]{
		'^': pair[int]{-1, 0},
		'>': pair[int]{0, 1},
		'v': pair[int]{1, 0},
		'<': pair[int]{0, -1},
	}

	for _, move := range moves {
		delta := charToMove[move]

		if shift(grid, pos, delta) {
			pos = pos.sum(delta)
		}

	}

	ans := calcAns(grid)
	fmt.Println(ans)

}
