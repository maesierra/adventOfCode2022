package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day18 struct {
}

type Cube struct {
	x int
	y int
	z int
}

func (c1 Cube) CommonSides(c2 Cube) []int {
	sides := []int{}
	diffX := c1.x - c2.x
	diffZ := c1.z - c2.z
	diffY := c1.y - c2.y
	if diffX == 1 && diffY == 0 && diffZ == 0 {
		sides = append(sides, 3)
	}
	if diffX == -1 && diffY == 0 && diffZ == 0 {
		sides = append(sides, 4)
	}
	if diffY == 1 && diffX == 0 && diffZ == 0 {
		sides = append(sides, 2)
	}
	if diffY == -1 && diffX == 0 && diffZ == 0 {
		sides = append(sides, 5)
	}
	if diffZ == 1 && diffY == 0 && diffX == 0 {
		sides = append(sides, 1)
	}
	if diffZ == -1 && diffY == 0 && diffX == 0 {
		sides = append(sides, 6)
	}
	return sides
}

func (c Cube) Key() string {
	return CoordsKey(c.x, c.y, c.z)
}

func (c Cube) Neighbours() []Cube{
	return []Cube{
		{c.x - 1, c.y,     c.z    },
		{c.x + 1, c.y,     c.z    },
		{c.x,     c.y - 1, c.z    },
		{c.x,     c.y + 1, c.z    },
		{c.x,     c.y,     c.z - 1},
		{c.x,     c.y,     c.z + 1},
	}
}

func (c Cube) String() string{
	return fmt.Sprintf("% 4d,% 4d,% 4d", c.x, c.y, c.z)
}

func (d Day18) NewCube(line string) Cube {
	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return Cube{x, y, z}
}

func (d Day18) SolvePart1(inputFile string, data []string) string {
	
	cubes := []*Cube{}
	total := 0
	for _, line := range common.ReadFileIntoLines(inputFile) {
		cube := d.NewCube(line)
		total += 6
		for _, c:= range cubes {
			common := len(c.CommonSides(cube))
			total = total - (common * 2)
		}
		cubes = append(cubes, &cube)
	}
	return strconv.Itoa(total)

}

func CoordsKey(x, y, z int) string {
	return fmt.Sprintf("%d-%d-%d", x, y, z)
}

type Limits struct {
	minX int
	minY int
	minZ int
	maxX int
	maxY int
	maxZ int
}

func CreateLimits() *Limits {
	return &Limits{
		minX : math.MaxInt,
		minY : math.MaxInt,
		minZ : math.MaxInt,
		maxX : math.MinInt,
		maxY : math.MinInt,
		maxZ : math.MinInt,
	}
}

func (l *Limits) Expand(cube Cube) {
	l.minX = common.IntMin(l.minX, cube.x)
	l.minY = common.IntMin(l.minY, cube.y)
	l.minZ = common.IntMin(l.minZ, cube.z)
	l.maxX = common.IntMax(l.maxX, cube.x)
	l.maxY = common.IntMax(l.maxY, cube.y)
	l.maxZ = common.IntMax(l.maxZ, cube.z)
}


func (d Day18) CanEscape(c Cube, l Limits, cubes map[string]*Cube, visited map[string]bool) bool{
	canEscape := true
	for x := c.x + 1; x <= l.maxX; x++ {
		if d.ContainsCube(x, c.y, c.z, cubes) {
			canEscape = false
			break
		}
	}
	if canEscape {
		return true
	}
	canEscape = true
	for x := c.x - 1; x >= l.minX; x-- {
		if d.ContainsCube(x, c.y, c.z, cubes) {
			canEscape = false
			break
		}
	}
	if canEscape {
		return true
	}
	canEscape = true
	for y := c.y + 1; y <= l.maxY; y++ {
		if d.ContainsCube(c.x, y, c.z, cubes) {
			canEscape = false
			break
		}
	}
	if canEscape {
		return true
	}
	canEscape = true
	for y := c.y - 1; y >= l.minY; y-- {
		if d.ContainsCube(c.x, y, c.z, cubes) {
			canEscape = false
			break
		}
	}
	if canEscape {
		return true
	}
	canEscape = true
	for z := c.z + 1; z <= l.maxZ; z++ {
		if d.ContainsCube(c.x, c.y, z, cubes) {
			canEscape = false
			break
		}
	}
	if canEscape {
		return true
	}
	canEscape = true
	for z := c.z - 1; z >= l.minZ; z-- {
		if d.ContainsCube(c.x, c.y, z, cubes) {
			canEscape = false
			break
		}
	}
	if canEscape {
		return true
	}
	visited[c.Key()] = true
	for _, n := range c.Neighbours() {
		if _, present := visited[n.Key()] ; present {
			continue
		}
		if d.ContainsCube(n.x, n.y, n.z, cubes) {
			continue
		}
		if d.CanEscape(n, l, cubes, visited) {
			return true
		}
	}
	return false
}

func (d Day18) ContainsCube(x, y, z int, cubes map[string]*Cube) bool {
	_, present := cubes[CoordsKey(x, y, z)];
	return present
}


func (d Day18) SolvePart2(inputFile string, data []string) string {
	cubes := map[string]*Cube{}
	total := 0
	limits := CreateLimits()
	
	for _, line := range common.ReadFileIntoLines(inputFile) {
		cube := d.NewCube(line)
		total += 6
		limits.Expand(cube)
		for _, c:= range cubes {			
			common := len(c.CommonSides(cube))
			total = total - (common * 2)
		}
		cubes[cube.Key()] = &cube
	}
	//Locate the empty spots that are touching at least other cube
	empty := map[string]*Cube{}
	for z := limits.minZ + 1; z < limits.maxZ; z++ {
		for y := limits.minY + 1 ; y < limits.maxY ; y++ {
			for x := limits.minX + 1 ; x < limits.maxX ; x++ {
				if !d.ContainsCube(x, y, z, cubes) {
					cube := Cube{x, y, z}		
					//Check if can reach the outside
					if !d.CanEscape(cube, *limits, cubes, map[string]bool{}) {						
						for _, c := range cube.Neighbours() {						
							if _, present := cubes[c.Key()]; present {
								empty[CoordsKey(x, y, z)] = &cube
								total -= len(cube.CommonSides(c))
							}
						}
					}
						
				}
			}
		}
	}
	for _, c := range empty {
		fmt.Printf("newEmpty(%v,%v,%d)\n", c.x, c.y, c.z)
	}
	for _, c := range cubes {
		fmt.Printf("newBox(%v,%v,%d)\n", c.x, c.y, c.z)
	}
	
	return strconv.Itoa(total)
}
