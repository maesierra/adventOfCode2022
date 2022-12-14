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

type Line struct {
	x1 int
	y1 int 
	x2 int 
	y2 int
}

func (l Line) IsHorizontal() bool{
	return l.y1 == l.y2
}

func (l Line) IsVertical() bool{
	return l.x1 == l.x2
}

func (l Line) Draw() []Point {
	points := []Point{}
	direction := 1
	if l.IsHorizontal() {		
		if l.x2 < l.x1 {
			direction = -1
		}
		for x := l.x1; x != l.x2; x+= direction {
			points = append(points, Point{x, l.y1})	
		}	
		points = append(points, Point{l.x2, l.y2})	
		return points
	} else if l.IsVertical() {
		if l.y2 < l.y1 {
			direction = -1
		}
		for y := l.y1; y != l.y2; y+= direction {
			points = append(points, Point{l.x1, y})	
		}	
		points = append(points, Point{l.x2, l.y2})	
		return points
	}
	panic("line not straight")
}

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type CavernMap struct {
	m *mat.Dense 
	minX int
	maxY int
	floor int	
	startingPoint Point
}

func (c CavernMap) HasFloor() bool {
	return c.floor != -1
}

func (c *CavernMap) AddWall(p Point) {
	c.m.Set(p.y, p.x, 3)
	c.minX = common.IntMin(c.minX, p.x)
	c.maxY = common.IntMax(c.maxY, p.y)
}


func (c CavernMap) IsSand(p Point) bool {
	return c.m.At(p.y, p.x) == 2 
}

func (c CavernMap) IsBlocked(p Point) bool{
	return c.m.At(p.y, p.x) != 0 
}

func (c * CavernMap) AddSand() *Point {
	point := Point{c.startingPoint.x, c.startingPoint.y}
	for ;; {
		if !c.HasFloor() && point.y > c.maxY {
			return nil //Falling into the abyss
		} else if c.HasFloor() && c.floor == point.y {
			c.m.Set(point.y - 1, point.x, 2)
			c.minX = common.IntMin(c.minX, point.x)
			return &Point{point.x, point.y - 1}
		}
		if c.IsBlocked(point) {
			if !c.IsBlocked(Point{point.x - 1, point.y}) {
				point.x-- 
				point.y++
			} else if !c.IsBlocked(Point{point.x + 1, point.y}) {
				point.x++ 
				point.y++
			} else {
				c.m.Set(point.y - 1, point.x, 2)
				c.minX = common.IntMin(c.minX, point.x)
				return &Point{point.x, point.y - 1}
			}
		} else {
			point.y++
		}
		if point.x == 0 {
			panic("Negative x")
		} else if point.x == c.m.RawMatrix().Cols - 1 {
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
	lines := []Line{}
	for _, line := range input {
		var last *Line = nil
		for _, l := range strings.Split(line, " -> ") {
			m := numbersRegExp.FindAllStringSubmatch(l, -1)
			x, _ := strconv.Atoi(m[0][0])
			y, _ := strconv.Atoi(m[1][0])
			maxX = common.IntMax(maxX, x)
			maxY = common.IntMax(maxY, y)
			if last != nil {
				last.x2 = x
				last.y2 = y				
				lines = append(lines, *last)
			} 
			newLine := Line{x, y, 0, 0}			
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
		Point{500, 0},
	}
	for _, line := range lines {
		for _, point := range line.Draw() {
			cavernMap.AddWall(point)
		}
	}
	return cavernMap
}

func (d Day14) SolvePart1(inputFile string) string {
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

func (d Day14) SolvePart2(inputFile string) string {
	cavernMap :=  d.ParseInput(inputFile, true)
	if debug {
		fmt.Printf("%v\n", cavernMap)	
	}
	nUnits := 0
	for nUnits = 1; ; nUnits++ {
		fmt.Printf("Units: %d\n", nUnits)
		sand := cavernMap.AddSand()
		fmt.Printf("Sand added at: %v\n", sand)
		if (sand.x == cavernMap.startingPoint.x && sand.y == cavernMap.startingPoint.y) {
			break;
		}
		if debug {
			fmt.Printf("%v\n", cavernMap)	
		}
	}
	return strconv.Itoa(nUnits)
}
