package main

import (
	"fmt"
	"os"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day1"
	"maesierra.net/advent-of-code/2022/adventofcode/day2"
	"maesierra.net/advent-of-code/2022/adventofcode/day3"
	"maesierra.net/advent-of-code/2022/adventofcode/day4"
	"maesierra.net/advent-of-code/2022/adventofcode/day5"
	"maesierra.net/advent-of-code/2022/adventofcode/day6"
)

func main() {

	args := os.Args[1:]
	if len(args) < 2 {
		println("Usage adventofcode <day> <part>")
		os.Exit(3)
	}
	day := os.Args[1]
	part := os.Args[2]
	var solver adventofcode.Solver
	switch day {
	case "1":
		solver = day1.Day1{}
	case "2":
		solver = day2.Day2{}
	case "3":
		solver = day3.Day3{}		
	case "4":
		solver = day4.Day4{}		
	case "5":
		solver = day5.Day5{}		
	case "6":
		solver = day6.Day6{}		
	default:
		solver = nil
	}

	if solver == nil {
		println("Day not implemented")
		os.Exit(3)
	}
	var result string
	inputFile := fmt.Sprintf("input%s", day)

	fmt.Printf("Running day %s part %s...\n", day, part)
	if part == "1" {
		result = solver.SolvePart1(inputFile)
	} else {
		result = solver.SolvePart2(inputFile)
	}
	fmt.Printf("Solution: %s\n", result)
}
