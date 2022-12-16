package day14

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
	"maesierra.net/advent-of-code/2022/common"
)

var debug = false

type Day14 struct {
}


type CavernMap struct {
	m *mat.Dense 
	minX int
	maxY int
	floor int	
	startingPoint common.Point
}

func (c CavernMap) HasFloor() bool {
	return c.floor != -1
}

func (c *CavernMap) AddWall(p common.Point) {
	c.m.Set(p.Y, p.X, 3)
	c.minX = common.IntMin(c.minX, p.X)
	c.maxY = common.IntMax(c.maxY, p.Y)
}


func (c CavernMap) IsSand(p common.Point) bool {
	return c.m.At(p.Y, p.X) == 2 
}

func (c CavernMap) IsBlocked(p common.Point) bool{
	return c.m.At(p.Y, p.X) != 0 
}

func (c * CavernMap) AddSand() *common.Point {
	point := common.Point{X: c.startingPoint.X, Y: c.startingPoint.Y}
	for ;; {
		if !c.HasFloor() && point.Y > c.maxY {
			return nil //Falling into the abyss
		} else if c.HasFloor() && c.floor == point.Y {
			c.m.Set(point.Y - 1, point.X, 2)
			c.minX = common.IntMin(c.minX, point.X)
			return &common.Point{X: point.X, Y: point.Y - 1}
		}
		if c.IsBlocked(point) {
			if !c.IsBlocked(common.Point{X: point.X - 1, Y: point.Y}) {
				point.X-- 
				point.Y++
			} else if !c.IsBlocked(common.Point{X: point.X + 1, Y: point.Y}) {
				point.X++ 
				point.Y++
			} else {
				c.m.Set(point.Y - 1, point.X, 2)
				c.minX = common.IntMin(c.minX, point.X)
				return &common.Point{X: point.X, Y: point.Y - 1}
			}
		} else {
			point.Y++
		}
		if point.X == 0 {
			panic("Negative x")
		} else if point.X == c.m.RawMatrix().Cols - 1 {
			c.m = mat.DenseCopyOf(c.m.Grow(0, 20))
		}
	}
}

func (c CavernMap) String() string {
	str := ""
	nCols := c.m.RawMatrix().Cols
	nRows := c.m.RawMatrix().Rows
	for r := 0; r < nRows; r++ {
		if r != 0 {
			str += "\n"
		}
		row := c.m.RowView(r)
		for col := c.minX - 2; col < nCols; col++ {
			if c.HasFloor() && r == c.floor {
				str += "#"
			} else {
				switch row.AtVec(col) {
				case 0:
					str += "."
				case 2: 
					str += "o"
				case 3:
					str += "#"	
				}	
			}
		}		
	}
	return str
}

func (d Day14) ParseInput(inputFile string, hasFloor bool) CavernMap {
	input := common.ReadFileIntoLines(inputFile)
	maxX := 0
	maxY := 0
	numbersRegExp, _ := regexp.Compile(`\d+`) 
	lines := []common.Line{}
	for _, line := range input {
		var last *common.Line = nil
		for _, l := range strings.Split(line, " -> ") {
			m := numbersRegExp.FindAllStringSubmatch(l, -1)
			x, _ := strconv.Atoi(m[0][0])
			y, _ := strconv.Atoi(m[1][0])
			maxX = common.IntMax(maxX, x)
			maxY = common.IntMax(maxY, y)
			if last != nil {
				last.X2 = x
				last.Y2 = y				
				lines = append(lines, *last)
			} 
			newLine := common.Line{X1: x, Y1: y, X2: 0, Y2: 0}			
			last = &newLine
		}
	}
	data := []float64{}
	for i := 0; i < maxY + 3; i++ {
		for j := 0; j < maxX + 3; j++ {
			data = append(data, 0)
		}
	}
	floor := -1
	if hasFloor {
		floor = maxY + 2
	}
	cavernMap := CavernMap{
		mat.NewDense(maxY + 3, maxX + 3, data), 
		math.MaxInt,
		0,
		floor,
		common.Point{X: 500, Y: 0},
	}
	for _, line := range lines {
		for _, point := range line.Draw() {
			cavernMap.AddWall(point)
		}
	}
	return cavernMap
}

func (d Day14) SolvePart1(inputFile string, data []string) string {
	cavernMap :=  d.ParseInput(inputFile, false)
	if debug {
		fmt.Printf("%v\n", cavernMap)	
	}
	nUnits := 0
	for nUnits = 1; cavernMap.AddSand() != nil; nUnits++ {
		fmt.Printf("Units: %d\n", nUnits)
		if debug {
			fmt.Printf("%v\n", cavernMap)	
		}
	}
	return strconv.Itoa(nUnits - 1)
}

func (d Day14) SolvePart2(inputFile string, data []string) string {
	cavernMap :=  d.ParseInput(inputFile, true)
	if debug {
		fmt.Printf("%v\n", cavernMap)	
	}
	nUnits := 0
	for nUnits = 1; ; nUnits++ {
		fmt.Printf("Units: %d\n", nUnits)
		sand := cavernMap.AddSand()
		fmt.Printf("Sand added at: %v\n", sand)
		if (sand.X == cavernMap.startingPoint.X && sand.Y == cavernMap.startingPoint.Y) {
			break;
		}
		if debug {
			fmt.Printf("%v\n", cavernMap)	
		}
	}
	return strconv.Itoa(nUnits)
}
