package common

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