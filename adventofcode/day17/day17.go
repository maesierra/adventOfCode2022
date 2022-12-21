package day17

import (
	"fmt"
	"math"
	"strconv"

	"maesierra.net/advent-of-code/2022/common"
)

type Day17 struct {

}

var debugLevel int = 0

type Block struct {
	line int
	size int
}
type Chamber struct {
	width int
	lastShape int 
	totalRocks int
	rocks map[int]*Rock
	perLine map[int]map[int]*Rock
	points map[int]map[int]bool
	maxBlockLeft Block
	maxBlockRight Block
	minHeight int
	maxHeight int
}

func (c *Chamber) AddRock() *Rock{
	c.lastShape = (c.lastShape + 1) % 5
	c.totalRocks++
	rock := Rock{n: c.totalRocks, shape: c.lastShape, placed: false}
	rock.Place(c.maxHeight + 4, c.width)
	c.rocks[rock.n] = &rock
	return &rock
}

func (c Chamber) MaxHeight() int{
	maxHeight := -1
	for _, r := range c.rocks {
		maxHeight = common.IntMax(maxHeight, r.box.Y2)
	}
	return maxHeight
}

func (c Chamber) MinHeight() int{
	min := math.MaxInt
	for _, r := range c.rocks {
		min = common.IntMin(min, r.box.Y1)		
	}
	return min
}

func (c *Chamber) MoveDown(rock *Rock) {
	if !c.move(rock, 0, -1) {
		rock.placed = true
	} 	
}

func (c *Chamber) MoveLeft(rock *Rock) {
	c.move(rock, -1, 0)
}

func (c *Chamber) MoveRight(rock *Rock) {
	c.move(rock, 1, 0)
}


func (c *Chamber) move(rock *Rock, x, y int) bool{
	if common.IntAbs(x) > 1 || common.IntAbs(y) > 1 {
		panic("Only movements of 1 are allowed")
	}
	box := rock.box.Move(x, y)
	// Check chamber boundaries
	if box.X1 < 0 || box.X2 >= c.width || box.Y1 < 0 {
		return false
	}
	hasCollisions := false
	for l := box.Y1; l <= box.Y2; l++ {
		if rocks, present := c.perLine[l] ; present {
			if len(rocks) == 1 {
				_, present = rocks[rock.n]
				if present {
					continue
				}
			}
			hasCollisions = true
			break
		}
	}
	if hasCollisions {
		for _, p := range rock.At(box) {
			row, present := c.points[p.Y] 
			if present {
				if _, present = row[p.X] ; present {
					return false
				}
			} 
			
	
			}
	}
	rock.box = box
	return true
}

func (c Chamber) Encode(r *Rock, position int) string {
	str := fmt.Sprintf("%v-%v-", r.shape, position)
	max := c.maxHeight
	min := c.minHeight
	for y := min; y < max; y++ {
		block := 0
		for x := 0; x < c.width; x++ {
			if _, present := c.points[y][x] ; present {
				str += "."
			} else {
				str += " "
			}
			
		}
		if block != 0 {
			str += strconv.Itoa(block)
		}
		str += "|"
	}
	return str
}


func (c Chamber) String() string {
	str := ""
	max := c.maxHeight
	min := c.minHeight
	for y := max; y >= min; y-- {
		str += fmt.Sprintf("% 6d |", y)
		for x := 0; x < c.width; x++ {
			cell := "."
			for _, rock := range c.perLine[y] {
				if rock.Contains(x, y) {
					if rock.placed {
						cell = "#"
					} else {
						cell = "@"
					}
					break
				}
			}
			str += cell
		}
		str += "|\n"
	
	}
	str += "       +"
	for x := 0; x < c.width; x++ { 
		str += "-"
	}
	return str + "+\n"
}

type Rock struct {
	n int
	shape int
	box common.BoundingBox
	placed bool
}

func (r *Rock) Place(atHeight, width int) {
	r.box.X1 = 2
	r.box.Y1 = atHeight
	switch r.shape {
	case 0:
		// #####
		r.box.X2 = r.box.X1 + 3
		r.box.Y2 = r.box.Y1
	case 1:
		// #
		//###
		// # 	
		r.box.X2 = r.box.X1 + 2
		r.box.Y2 = r.box.Y1 + 2
	case 2:
		//  #
		//  #
		//###	
		r.box.X2 = r.box.X1 + 2
		r.box.Y2 = r.box.Y1 + 2
	case 3:
		//  #
		//  #
		//  #	
		//  #
		r.box.X2 = r.box.X1
		r.box.Y2 = r.box.Y1 + 3
	case 4:
		//  ##
		//  ##
		r.box.X2 = r.box.X1 + 1
		r.box.Y2 = r.box.Y1 + 1
	}
}

func (r1 Rock) Intersects(r2 Rock, at common.BoundingBox) bool {
	if !r1.box.Intersects(at) {
		return false
	}
	for _, p := range r2.At(at) {
		if r1.Contains(p.X, p.Y) {
			return true
		}
	}
	return false
}
func (r Rock) Shape() [][]bool {
	switch r.shape {
	case 0:
		// #####
		return [][]bool{
			{true, true, true, true},
		}
	case 1:
		// #
		//###
		// # 	
		return [][]bool{
			{false, true, false},
			{true,  true, true},
			{false, true, false},
		}
	case 2:
		//  #
		//  #
		//###	
		return [][]bool{
			{false, false, true},
			{false, false, true},
			{true , true , true},
		}
	case 3:
		//  #
		//  #
		//  #	
		//  #
		return [][]bool{
			{true},
			{true},
			{true},
			{true},
		}
	default:
		//  ##
		//  ##
		return [][]bool{
			{true, true},
			{true, true},
		}
	}
}

func (r Rock) CountPoints() int {
	switch r.shape {
	case 0:
		// #####
		return 4
	case 1:
		// #
		//###
		// # 	
		return 5
	case 2:
		//  #
		//  #
		//###	
		return 5
	case 3:
		//  #
		//  #
		//  #	
		//  #
		return 4
	default:
		//  ##
		//  ##
		return 4
	}
}

func (r Rock) Contains(x, y  int) bool {
	if !r.box.Contains(x, y) {
		return false
	}
	shape := r.Shape()
	l := len(shape) -1 
	return shape[l - (y - r.box.Y1)][x - r.box.X1]
}

func (r Rock) At(box common.BoundingBox) []common.Point {
	shape := r.Shape()
	l := len(shape) - 1
	points := make([]common.Point, r.CountPoints())
	c := 0
	for y := box.Y1; y <= box.Y2; y++ {
		for x := box.X1; x <= box.X2; x++ {
			if shape[l - (y - box.Y1)][x - box.X1] {
				points[c] = common.Point{X: x, Y: y}
				c++				
			}
		}
	}
	return points
}

func (c *Chamber) CalculateMaxBlockLeft(from, to int) {
	for y:= from ; y <= to ; y++  {
		continuous := true
		for x := 0; continuous && x < c.width; x++ { 
			continuous = false
			for _, r := range c.perLine[y] {
				if r.Contains(x, y) {
					continuous = true
					break
				}
			}
			if continuous {
				if c.maxBlockLeft.size <= x + 1 {
					c.maxBlockLeft.line = y
					c.maxBlockLeft.size = x + 1
				}	
			}
		}
	}
} 
func (c *Chamber) CalculateMaxBlockRight(from, to int) {
	for y:= from ; y <= to ; y++  {
		continuous := true
		for x := c.width - 1; continuous && x >= 0; x-- { 
			continuous = false
			for _, r := range c.perLine[y] {
				if r.Contains(x, y) {
					continuous = true
					break
				}
			}
			if continuous {
				if c.maxBlockRight.size <= c.width - x {
					c.maxBlockRight.line = y
					c.maxBlockRight.size = c.width - x
				}	
			}
		}
	}
} 
func (c *Chamber) Prune(rock *Rock) {
	//Determine if by placing this rock a line will be blocked
	c.CalculateMaxBlockLeft(rock.box.Y1, rock.box.Y2)
	c.CalculateMaxBlockRight(rock.box.Y1, rock.box.Y2)
	if c.maxBlockLeft.size + c.maxBlockRight.size > c.width {			
		count := 0
		//prune rocks that cannot affect anymore
		minLine := c.maxBlockLeft.line
		if c.maxBlockRight.line < c.maxBlockLeft.line {
			minLine = c.maxBlockRight.line
		}
		linesToDelete := []int{}
		rocksToDelete := map[int]bool{}
		for l, rocks := range c.perLine {
			if l < minLine  {
				toDeleteInLine := 0		
				for _, r := range rocks {
					if r.box.Y2 < minLine {
						rocksToDelete[r.n] = true
						toDeleteInLine++
						count++
					}
				}
				if toDeleteInLine == len(rocks) {
					linesToDelete = append(linesToDelete, l)
				} 
				
			}
		}
		for n, _ := range rocksToDelete {
			rock := c.rocks[n]
			for y := rock.box.Y1; y <= rock.box.Y2; y++ {
				delete(c.perLine[y], n)
			}
			delete(c.rocks, n)
			for _, p := range rock.At(rock.box) {
				delete(c.points[p.Y], p.X)
				if len(c.points[p.Y]) == 0 {
					delete(c.points, p.Y)
				}
			}
		}
		for _, l :=  range linesToDelete {
			delete(c.perLine, l)
		}
		if debugLevel > 1 {
			fmt.Printf("Pruned %d items\n", count)
		}
		if debugLevel > 3 {
			fmt.Printf("%v", c)
		}
		if minLine == c.maxBlockLeft.line {
			c.maxBlockLeft.size = 0
			c.CalculateMaxBlockLeft(minLine + 1, c.maxBlockRight.line)
		} else {
			c.maxBlockRight.size = 0
			c.CalculateMaxBlockRight(minLine + 1, c.maxBlockLeft.line)
		}
	}
}
func (d Day17) RunSimulation(movements string, nRocks uint64) int {
	pos := 0
	c := Chamber{
		lastShape: 4, 
		width: 7, 
		rocks: map[int]*Rock{}, 
		perLine: map[int]map[int]*Rock{},
		points: map[int]map[int]bool{},
		maxHeight: -1,
		totalRocks: 0,
	}
	cache := map[string]int{}
	var i uint64 
	nMovements := len(movements)
	cycleStarted := ""
	cycle :=[]int{}
	for i = 0; i < nRocks; i++ {		
		rock := c.AddRock();
		if debugLevel > 0 {
			fmt.Printf("Adding rock %d min: %d max: %d nRocks: %d\n", i + 1, c.minHeight, c.maxHeight, len(c.rocks))
		} else if debugLevel ==0 && i % 1000000 ==0 {
			fmt.Printf("Adding rock %d nRocks: %d\n", i + 1, len(c.rocks))
		}
		for ; !rock.placed ;  {
			if movements[pos] == '<' {
				c.MoveLeft(rock)
			} else {
				c.MoveRight(rock)
			}
			c.MoveDown(rock)
			pos = (pos + 1) % nMovements
		}
		for _, p := range rock.At(rock.box) {
			_, present := c.points[p.Y] 
			if !present {
				c.points[p.Y] = map[int]bool{p.X: true}
			} else {
				c.points[p.Y][p.X] = true
			}
		}
		for y := rock.box.Y1; y <= rock.box.Y2; y++ {
			_, present := c.perLine[y] 
			if !present {
				c.perLine[y] = map[int]*Rock{rock.n: rock}
			} else {
				c.perLine[y][rock.n] = rock
			}
		}
		c.minHeight = c.MinHeight()
		c.maxHeight = c.MaxHeight()
		key := c.Encode(rock, pos)
		if cached, present := cache[key]; present {
			if cycleStarted == "" {
				cycleStarted = key
				fmt.Printf("Cache hit after %d %v %v %v\n", i, c.maxHeight, cached, c.maxHeight - cached)
				cycle = append(cycle, c.maxHeight)
			} else if cycleStarted == key {
				linesPerCycle := cycle[len(cycle) - 1] - cycle[0]
				nTimes := nRocks / uint64(len(cycle))
				fmt.Printf("%d %v other %v\n", linesPerCycle, len(cycle), c.maxHeight - cycle[1])
				total := nTimes * uint64(linesPerCycle) + uint64(cycle[int(nRocks % uint64(len(cycle)))] - cycle[0])
				return int(total)
			}  else {
				cycle = append(cycle, c.maxHeight)
			}
		} else {
			cache[key] = c.maxHeight
		}
		if debugLevel > 3 {
			fmt.Printf("%v", c)
		}
		c.Prune(rock)
	}	
	if debugLevel > 2 {
		fmt.Printf("%v", c)
	}
	return c.MaxHeight() + 1
}

func (d Day17) SolvePart1(inputFile string, data []string) string {	
	input := common.ReadFile(inputFile)
	return strconv.Itoa(d.RunSimulation(input, 2022))

}

func (d Day17) SolvePart2(inputFile string, data []string) string {
	input := common.ReadFile(inputFile)
	return strconv.Itoa(d.RunSimulation(input, 1000000000000))
}
