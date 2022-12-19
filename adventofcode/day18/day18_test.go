package day18_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day18"
)

var solver adventofcode.Solver = day18.Day18{}

func Test_SolvePart1(t *testing.T) {
	input := `1,1,1
2,1,1`
	expected := "10"
	adventofcode.SolvePart1Test(input, expected, solver, t)

	input = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`
	expected = "64"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`
	expected := "58"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
