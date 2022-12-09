package common

import "math"

// Max returns the larger of x or y.
func IntMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func IntMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func IntAbs(x int) int {
	return int(math.Abs(float64(x)))
}