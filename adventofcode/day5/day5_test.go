package day5_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day5"
)

var solver adventofcode.Solver = day5.Day5{}

func Test_SolvePart1(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	expected := "CMZ"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart1_SecondInput(t *testing.T) {
	input := `[S]                 [T] [Q]
[L]             [B] [M] [P]     [T]
[F]     [S]     [Z] [N] [S]     [R]
[Z] [R] [N]     [R] [D] [F]     [V]
[D] [Z] [H] [J] [W] [G] [W]     [G]
[B] [M] [C] [F] [H] [Z] [N] [R] [L]
[R] [B] [L] [C] [G] [J] [L] [Z] [C]
[H] [T] [Z] [S] [P] [V] [G] [M] [M]
 1   2   3   4   5   6   7   8   9`
	expected := "SRSJBTQRT"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}


func Test_SolvePart2(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	expected := "MCD"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
