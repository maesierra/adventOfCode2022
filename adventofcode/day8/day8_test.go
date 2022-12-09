package day8_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day8"
)

var solver adventofcode.Solver = day8.Day8{}

func Test_SolvePart1(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	expected := "21"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	expected := "8"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
