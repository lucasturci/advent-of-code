package main

import (
	"errors"
	"fmt"
)

const costA = int64(3)
const costB = int64(1)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func extendedGCDImpl(a, b int64, ra, rb []int64) (int64, int64) {
	if b == 0 {
		return ra[0], ra[1]
	}

	k := a / b
	return extendedGCDImpl(b, a%b, rb, []int64{ra[0] - k*rb[0], ra[1] - k*rb[1]})
}

// finds a linear combination of (a, b) that represents its gcd
func extendedGCD(a, b int64) (int64, int64) {
	ra := []int64{1, 0}
	rb := []int64{0, 1}
	return extendedGCDImpl(a, b, ra, rb)
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func diophantine(a, b, c int64) (int64, int64, error) {
	x, y := extendedGCD(a, b)
	g := gcd(a, b)

	if c%g != 0 {
		return 0, 0, errors.New("no solution")
	}
	return x * c / g, y * c / g, nil
}

// x * a[0] + y * b[0] = at[0]
// x * a[1] + y * b[1] = at[1]
// min x * 3 + y
func solve(a, b, at []int64) int64 {
	det := a[0]*b[1] - a[1]*b[0]
	at[0] += 10000000000000
	at[1] += 10000000000000
	if det == 0 { // treat the case of infinite solutions or zero solutions
		if at[1]*b[0] == at[0]*b[1] { // means there are infint64e solutions
			// take any equation
			x, y, err := diophantine(a[0], b[0], at[0])
			if err != nil {
				return 0
			}

			lcm := abs(x * y / gcd(x, y))
			if y < 0 {
				x, y = -x, -y
			}
			incx := lcm / x
			incy := lcm / y
			if x < 0 {
				// x + k * incx > 0 => k > -x/incx
				k := (-x + incx - 1) / incx
				x += k * incx
				y -= k * incy
			}
			if x < 0 || y < 0 {
				return 0
			}

			// now minimize x. x >= 0 and y >=  0
			// x - k * incx >= 0 => k >= x/incx
			k := x / incx
			x -= k * incx
			y += k * incy

			return 3*x + y
		} else {
			return 0
		}
	}
	dx := at[0]*b[1] - at[1]*b[0]
	dy := a[0]*at[1] - a[1]*at[0]
	if dx%det != 0 || dy%det != 0 || dx/det < 0 || dy/det < 0 {
		return 0
	}
	x, y := dx/det, dy/det

	return x*costA + y*costB
}

// 1 2
// 2 4
//
//
//
//
//
//

func main() {
	var ans int64
	for {
		a := make([]int64, 2)
		b := make([]int64, 2)
		at := make([]int64, 2)
		var read int
		read, _ = fmt.Scanf("Button A: X+%d, Y+%d\n", &a[0], &a[1])
		if read != 2 {
			break
		}
		read, _ = fmt.Scanf("Button B: X+%d, Y+%d\n", &b[0], &b[1])
		if read != 2 {
			break
		}
		read, _ = fmt.Scanf("Prize: X=%d, Y=%d\n", &at[0], &at[1])
		if read != 2 {
			break
		}
		ans += solve(a, b, at)
		fmt.Scanf("\n")
	}

	fmt.Println(ans)
}
