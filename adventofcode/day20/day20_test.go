package day20_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day20"
)

var solver adventofcode.Solver = day20.Day20{}

func Test_SolvePart1(t *testing.T) {
	input := `1
2
-3
3
-2
0
4`
	expected := "3"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `1
2
-3
3
-2
0
4`
	expected := "1623178306"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
