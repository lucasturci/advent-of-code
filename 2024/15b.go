package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

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

		// do the replaces specified in the statement
		line = strings.Replace(line, "#", "##", -1)
		line = strings.Replace(line, "O", "[]", -1)
		line = strings.Replace(line, ".", "..", -1)
		line = strings.Replace(line, "@", "@.", -1)

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
// For part 2, let's apply the same rules, but make each side stick with the other
// This will incur in more than one branch, but should be at most size of grid calls
// We have to apply moves after
// This ideally should have a vis map, but input doesn't require one for complexity improvement
func shift(grid [][]byte, pos pair[int], delta pair[int]) (bool, [][]pair[int]) {
	// fmt.Println("DELTA", delta)
	newPos := pos.sum(delta)

	moves := [][]pair[int]{}

	if grid[pos[0]][pos[1]] == '#' {
		return false, moves
	}

	if grid[pos[0]][pos[1]] == '.' {
		return true, moves
	}

	move := func(pos pair[int], delta pair[int]) {
		moves = append(moves, []pair[int]{pos, delta})
	}

	if delta[1] != 0 { // moving right or left is the same.
		ret, mov := shift(grid, newPos, delta)
		if ret {
			moves = append(mov, moves...)
			move(pos, delta)
		}
		return ret, moves
	}
	// panic("should not arrive here")

	c := grid[pos[0]][pos[1]]
	if slices.Contains([]byte{'[', ']'}, c) {
		var other pair[int]
		if c == '[' {
			other = pos.sum(pair[int]{0, 1})
		} else {
			other = pos.sum(pair[int]{0, -1})
		}
		c2 := grid[newPos[0]][newPos[1]]
		ret1, mov1 := shift(grid, pos.sum(delta), delta)
		mov2 := [][]pair[int]{}
		ret2 := true
		if c2 != c { // other case means that blocks are aligned, so only call one.
			ret2, mov2 = shift(grid, other.sum(delta), delta)
			mov1 = append(mov1, mov2...)
		}
		ret := ret1 && ret2
		if ret {
			moves = append(mov1, moves...)
			move(pos, delta)
			move(other, delta)
		}
		return ret, moves
	}

	if ok, mov := shift(grid, newPos, delta); ok { // for case of @
		moves = append(mov, moves...)
		move(pos, delta)
		return true, moves
	}

	return false, [][]pair[int]{}

}

func calcAns(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '[' {
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

		if ok, shifts := shift(grid, pos, delta); ok {
			pos = pos.sum(delta)
			// process moves
			found := make(map[int]bool)
			for _, s := range shifts {
				pos, delta := s[0], s[1]

				// discard repeated moves
				if found[1000*pos[0]+pos[1]] {
					continue
				}
				found[1000*pos[0]+pos[1]] = true

				newPos := pos.sum(delta)
				grid[newPos[0]][newPos[1]] = grid[pos[0]][pos[1]]
				grid[pos[0]][pos[1]] = '.'
			}
		}

	}
	printGrid(grid)

	ans := calcAns(grid)
	fmt.Println(ans)

}
