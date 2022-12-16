package adventofcode

type Solver interface {
	SolvePart1(inputFile string, data []string) string
	SolvePart2(inputFile string, data []string) string
}
