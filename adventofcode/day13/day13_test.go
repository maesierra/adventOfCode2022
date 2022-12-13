package day13_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day13"
)

var solver adventofcode.Solver = day13.Day13{}

func Test_SolvePart1(t *testing.T) {
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`
	expected := "13"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`
	expected := "140"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
