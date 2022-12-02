package day2_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day2"
)

var solver adventofcode.Solver = day2.Day2{}

func Test_SolvePart1(t *testing.T) {
	input := `A Y
B X
C Z`
	expected := "15"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `A Y
B X
C Z`
	expected := "12"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
