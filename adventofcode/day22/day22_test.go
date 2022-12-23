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


func FaceTest(dir string, fromX, fromY, n int, newDir string, toX, toY int, b *day22.Board, t *testing.T) {
	b.Position.X = fromX
	b.Position.Y = fromY
	b.Direction = dir
	b.Move(day22.Movement{N: n, Rotation: ""})
	if b.Position.X != toX || b.Position.Y != toY {
		t.Errorf("FacesTest() = %q, want %q", b.Position, common.Point{X: toX, Y: toY})
	}
	if b.Direction != newDir {
		t.Errorf("FacesTest() = %q, want %q", b.Direction, newDir)
	}
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
	//Face 1T to 2T
	FaceTest("^", 10, 0, 3, "v", 1, 6, &board, t)
	//Face 1R to 6T
	FaceTest(">", 10, 1, 3, "<", 14, 10, &board, t)
	//Face 1B to 4T
	FaceTest("v", 9, 3, 4, "v", 9, 7, &board, t)
	//Face 1L to 3T
	FaceTest("<", 8, 1, 2, "v", 5, 5, &board, t)

	//Face 2T to 1T
	FaceTest("^", 1, 7, 5, "v", 10, 1, &board, t)
	//Face 2R to 3L
	FaceTest(">", 2, 5, 3, ">", 5, 5, &board, t)
	//Face 2B to 5B
	FaceTest("v", 2, 6, 3, "^", 9, 11, &board, t)
	//Face 2L to 6B
	FaceTest("<", 1, 6, 6, "^", 13, 10, &board, t)
}
