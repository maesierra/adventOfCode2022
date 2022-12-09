package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
	"maesierra.net/advent-of-code/2022/common"
	"maesierra.net/advent-of-code/2022/common/intmatrix"
)

type Day9 struct {
}

type Movement struct {
	direction string
	steps     int
}

type Status float64

func (d Day9) ParseFile(inputFile string) []Movement {
	res := []Movement{}
	for _, line := range common.ReadFileIntoLines(inputFile) {
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[1])
		res = append(res, Movement{parts[0], steps})
	}
	return res
}

type Position struct {
	row    int
	column int
}

func (p1 Position) Distance(p2 Position) int {
	return int(math.Max(math.Abs(float64(p1.column-p2.column)), math.Abs(float64(p1.row-p2.row))))
}

type RopeBridge struct {
	size int
	sections []Position
	headPosition Position
	bridgeMap    *mat.Dense
	visited      *mat.Dense
	countVisited int
}
func (b *RopeBridge) Set(p Position, value float64) {
	b.bridgeMap.Set(p.row, p.column, value)
}


func (b *RopeBridge) MoveHead(mov Movement) {
	fmt.Printf("Move %s %d\n", mov.direction, mov.steps)
	for n := 0; n < mov.steps; n++ {
		//Clean up before the movement
		b.Set(b.headPosition, 0)
		for _, p := range b.sections {
			b.Set(p, 0)
		}

		switch mov.direction {
		case "R":
			b.headPosition.column++
		case "L":
			b.headPosition.column--
		case "D":
			b.headPosition.row++
		case "U":
			b.headPosition.row--
		}
		if b.headPosition.column >= b.bridgeMap.RawMatrix().Cols {
			b.bridgeMap = intmatrix.AddColumnToTheRight(b.bridgeMap)
			b.visited = intmatrix.AddColumnToTheRight(b.visited)
		} else if b.headPosition.row >= b.bridgeMap.RawMatrix().Rows {
			b.bridgeMap = intmatrix.AddRowToTheBottom(b.bridgeMap)
			b.visited = intmatrix.AddRowToTheBottom(b.visited)
		} else if b.headPosition.row < 0 {
			b.bridgeMap = intmatrix.AddRowToTheTop(b.bridgeMap)
			b.visited = intmatrix.AddRowToTheTop(b.visited)
			b.headPosition.row = 0
			for idx := range b.sections {
				b.sections[idx].row++
			}
		} else if b.headPosition.column < 0 {
			b.bridgeMap = intmatrix.AddColumnToTheLeft(b.bridgeMap)
			b.visited = intmatrix.AddColumnToTheLeft(b.visited)
			b.headPosition.column = 0
			for idx := range b.sections {
				b.sections[idx].column++
			}
		}
		
		//Move the sections		
		for i := len(b.sections) - 1; i >= 0; i-- {
			var head * Position
			if (i == len(b.sections) - 1) {
				head = &b.headPosition
			} else {
				head = &b.sections[i + 1]
			}
			distance := head.Distance(b.sections[i])
			if distance > 1 {
				if head.column < b.sections[i].column {
					b.sections[i].column--
				} else if head.column > b.sections[i].column {
					b.sections[i].column++
				}

				if head.row < b.sections[i].row {
					b.sections[i].row--
				} else if head.row > b.sections[i].row {
					b.sections[i].row++
				}
				if i == 0 {
					alreadyVisited := b.visited.At(b.sections[i].row, b.sections[i].column)
					if alreadyVisited == 0 {
						b.visited.Set(b.sections[i].row, b.sections[i].column, 1)
						b.countVisited++
					}	
				}
			}
			b.Set(b.sections[i], float64(b.size - i))
		}		
		b.Set(b.headPosition, float64(b.size + 1))
		// fmt.Println(mat.Formatted(b.bridgeMap))
		// fmt.Println()
	}
}

func (d Day9) SolvePart1(inputFile string) string {

	//starts in the bottom possition for a 2x2 grid
	bridge := RopeBridge{
		1,
		[]Position{{1, 0}},
		Position{1, 0},
		mat.NewDense(2, 2, []float64{0, 0, 2, 0}),
		mat.NewDense(2, 2, []float64{0, 0, 1, 0}),
		1,
	}
	for _, movement := range d.ParseFile(inputFile) {
		bridge.MoveHead(movement)
	}
	fmt.Println(mat.Formatted(bridge.visited))
	return strconv.Itoa(bridge.countVisited)

}

func (d Day9) SolvePart2(inputFile string) string {
	//starts in the bottom possition for a 2x2 grid
	bridge := RopeBridge{
		9,
		[]Position{{1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}},
		Position{1, 0},
		mat.NewDense(2, 2, []float64{0, 0, 10, 0}),
		mat.NewDense(2, 2, []float64{0, 0, 1, 0}),
		1,
	}
	for _, movement := range d.ParseFile(inputFile) {
		bridge.MoveHead(movement)
	}
	fmt.Println(mat.Formatted(bridge.visited))
	return strconv.Itoa(bridge.countVisited)
}
