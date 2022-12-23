package day22_test

import (
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day22"
	"maesierra.net/advent-of-code/2022/common"
)

var solver adventofcode.Solver = day22.Day22{}

func Test_SolvePart1(t *testing.T) {
	input := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	expected := "6032"
	adventofcode.SolvePart1WithDataTest(input, []string{"1"}, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	expected := "5031"
	adventofcode.SolvePart2WithDataTest(input, []string{"1"}, expected, solver, t)
}

func Test_3DMovements(t *testing.T) {
	input := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	d := day22.Day22{}
	m := day22.Maps3d[1]
	board, _ := d.ParseInput(adventofcode.StringToInputFile(input), &m)
	//Face 1R to 6T
	board.Position.X = 10
	board.Position.Y = 1
	board.Move(day22.Movement{N: 3, Rotation: ""})
	if board.Position.X != 14 || board.Position.Y != 9 {
		t.Errorf("Part2() = %q, want %q", board.Position, common.Point{X: 14, Y: 9})
	}
}
