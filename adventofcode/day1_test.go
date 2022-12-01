package adventofcode_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
)

var solver adventofcode.Solver = adventofcode.Day1{}

func Test_SolvePart1(t *testing.T) {
	input := "aaabbcc"
	expected := "d5bda34132f7b4b674b49e44c79b3f95"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := "aaabbcc"
	expected := "6fba23697ff86022e4895b61a4e141f6"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
