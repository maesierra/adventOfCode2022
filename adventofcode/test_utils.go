package adventofcode

import (
	"testing"

	"maesierra.net/advent-of-code/2022/common"
)

func StringToInputFile(contents string) string {
	file := common.CreateTempFile(contents)
	return file.Name()
}

func SolvePart1Test(input string, expectedOutput string, s Solver, t *testing.T) {
	SolvePart1WithDataTest(input, []string{}, expectedOutput, s, t)
}

func SolvePart1WithDataTest(input string, data []string, expectedOutput string, s Solver, t *testing.T) {
	got := s.SolvePart1(StringToInputFile(input), data)
	if got != expectedOutput {
		t.Errorf("Part1() = %q, want %q", got, expectedOutput)
	}
}

func SolvePart2Test(input string, expectedOutput string, s Solver, t *testing.T) {
	SolvePart2WithDataTest(input, []string{}, expectedOutput, s, t)

}

func SolvePart2WithDataTest(input string, data []string, expectedOutput string, s Solver, t *testing.T) {
	got := s.SolvePart2(StringToInputFile(input), data)
	if got != expectedOutput {
		t.Errorf("Part2() = %q, want %q", got, expectedOutput)
	}

}