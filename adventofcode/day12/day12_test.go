package day12_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day12"
)

var solver adventofcode.Solver = day12.Day12{}

func Test_SolvePart1(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	expected := "31"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	expected := "29"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
