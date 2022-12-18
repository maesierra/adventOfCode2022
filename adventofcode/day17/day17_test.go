package day17_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day17"
)

var solver adventofcode.Solver = day17.Day17{}

func Test_SolvePart1(t *testing.T) {
	input := `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`
	expected := "3068"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`
	expected := "1514285714288"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
