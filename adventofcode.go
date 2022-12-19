package main

import (
	"fmt"
	"os"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day1"
	"maesierra.net/advent-of-code/2022/adventofcode/day10"
	"maesierra.net/advent-of-code/2022/adventofcode/day11"
	"maesierra.net/advent-of-code/2022/adventofcode/day12"
	"maesierra.net/advent-of-code/2022/adventofcode/day13"
	"maesierra.net/advent-of-code/2022/adventofcode/day14"
	"maesierra.net/advent-of-code/2022/adventofcode/day15"
	"maesierra.net/advent-of-code/2022/adventofcode/day17"
	"maesierra.net/advent-of-code/2022/adventofcode/day18"
	"maesierra.net/advent-of-code/2022/adventofcode/day2"
	"maesierra.net/advent-of-code/2022/adventofcode/day3"
	"maesierra.net/advent-of-code/2022/adventofcode/day4"
	"maesierra.net/advent-of-code/2022/adventofcode/day5"
	"maesierra.net/advent-of-code/2022/adventofcode/day6"
	"maesierra.net/advent-of-code/2022/adventofcode/day7"
	"maesierra.net/advent-of-code/2022/adventofcode/day8"
	"maesierra.net/advent-of-code/2022/adventofcode/day9"
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
	case "7":
		solver = day7.Day7{}		
	case "8":
		solver = day8.Day8{}		
	case "9":
		solver = day9.Day9{}		
	case "10":
		solver = day10.Day10{}		
	case "11":
		solver = day11.Day11{}		
	case "12":
		solver = day12.Day12{}		
	case "13":
		solver = day13.Day13{}		
	case "14":
		solver = day14.Day14{}		
	case "15":
		solver = day15.Day15{}		
	case "17":
		solver = day17.Day17{}		
	case "18":
		solver = day18.Day18{}		
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
		result = solver.SolvePart1(inputFile, []string{})
	} else {
		result = solver.SolvePart2(inputFile, []string{})
	}
	fmt.Printf("Solution: %s\n", result)
}
