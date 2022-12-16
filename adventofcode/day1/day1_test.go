package day1_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day1"
)

var solver adventofcode.Solver = day1.Day1{}

func Test_SolvePart1(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	expected := "24000"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	expected := "45000"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
