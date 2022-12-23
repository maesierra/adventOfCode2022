package day23_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day23"
)

var solver adventofcode.Solver = day23.Day23{}

func Test_SolvePart1(t *testing.T) {
	input := `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`
	expected := "110"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`
	expected := "20"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
