package day14_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day14"
)

var solver adventofcode.Solver = day14.Day14{}

func Test_SolvePart1(t *testing.T) {
	input := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
	expected := "24"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
	expected := "93"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
