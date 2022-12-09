package day9_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day9"
)

var solver adventofcode.Solver = day9.Day9{}

func Test_SolvePart1(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	expected := "13"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	expected := "1"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}



func Test_SolvePart2_Sample2(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	expected := "36"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}