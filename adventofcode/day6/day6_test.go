package day6_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day6"
)

var solver adventofcode.Solver = day6.Day6{}

func Test_SolvePart1(t *testing.T) {
	input := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	expected := "7"
	adventofcode.SolvePart1Test(input, expected, solver, t)

	input = `bvwbjplbgvbhsrlpgdmjqwftvncz`
	expected = "5"
	adventofcode.SolvePart1Test(input, expected, solver, t)

	input = `nppdvjthqldpwncqszvftbrmjlhg`
	expected = "6"
	adventofcode.SolvePart1Test(input, expected, solver, t)

	input = `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`
	expected = "10"
	adventofcode.SolvePart1Test(input, expected, solver, t)
	
	input = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`
	expected = "11"
	adventofcode.SolvePart1Test(input, expected, solver, t)

}

func Test_SolvePart2(t *testing.T) {
	input := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	expected := "19"
	adventofcode.SolvePart2Test(input, expected, solver, t)

	input = `bvwbjplbgvbhsrlpgdmjqwftvncz`
	expected = "23"
	adventofcode.SolvePart2Test(input, expected, solver, t)

	input = `nppdvjthqldpwncqszvftbrmjlhg`
	expected = "23"
	adventofcode.SolvePart2Test(input, expected, solver, t)

	input = `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`
	expected = "29"
	adventofcode.SolvePart2Test(input, expected, solver, t)
	
	input = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`
	expected = "26"
	adventofcode.SolvePart2Test(input, expected, solver, t)

}
