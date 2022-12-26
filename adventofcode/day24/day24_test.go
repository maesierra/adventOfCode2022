package day24_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day24"
)

var solver adventofcode.Solver = day24.Day24{}

func Test_SolvePart1(t *testing.T) {
	input := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
	expected := "18"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
	expected := "54"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
