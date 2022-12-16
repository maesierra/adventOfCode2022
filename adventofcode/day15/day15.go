package day15

import (
	"fmt"
	"regexp"
	"strconv"

	"maesierra.net/advent-of-code/2022/common"
)

var debug bool = false
type Day15 struct {
}

type Sensor struct {
	position common.Point
	closestBeacon common.Point
}

func (s Sensor) Distance() int {
	return s.position.ManhattanDistance(s.closestBeacon)
}


func (s Sensor) String() string {
	return fmt.Sprintf("Sensor at %v", s.position)
}

type Grid struct {
	sensors []Sensor
	sensorPoints map[int]map[int]bool
	beaconPoints map[int]map[int]bool
}

func (g Grid) NItemsInRow(y int) int{
	return g.inRow(g.beaconPoints, y) + g.inRow(g.sensorPoints, y)
}

func (g Grid) inRow(aMap map[int]map[int]bool, y int) int {
	row, ok := aMap[y]
	if ok {
		return len(row)
	} else {
		return 0
	}
}

func (g Grid) add(aMap map[int]map[int]bool, p common.Point) {
	_, hasKey := aMap[p.Y]
	if hasKey {
		aMap[p.Y][p.X] = true
	} else {
		aMap[p.Y] = map[int]bool{p.X: true}
	}
}

func (g Grid) check(aMap map[int]map[int]bool, p common.Point) bool{
	_, hasY := aMap[p.Y]
	if hasY {
		_, hasX := aMap[p.Y][p.X]
		return hasX
	} else {
		return false
	}
}

func (g Grid) CheckPoint(p common.Point) bool {
	return g.check(g.sensorPoints, p) || g.check(g.beaconPoints, p)
}


func (g * Grid) InArea(y int, sensor Sensor) (common.Range, bool) {
	distance := sensor.Distance()
	distanceY := common.IntAbs(sensor.position.Y - y)
	startY := sensor.position.Y - distance 
	endY := sensor.position.Y + distance
	startX := sensor.position.X - distance + distanceY
	endX := sensor.position.X + distance - distanceY
	if y < startY || y > endY {
		return common.Range{}, false
	} else {
		return common.NewRange(startX, endX), true
	}

}


func (d Day15) CreateGrid(inputFile string) Grid {
	numbersRegExp, _ := regexp.Compile(`-?\d+`)
	sensors:=  []Sensor {}
	for _, line := range common.ReadFileIntoLines(inputFile) {
		m := numbersRegExp.FindAllStringSubmatch(line, -1)
		sensorX, _ := strconv.Atoi(m[0][0])
		sensorY, _ := strconv.Atoi(m[1][0])
		beaconX, _ := strconv.Atoi(m[2][0])
		beaconY, _ := strconv.Atoi(m[3][0])
		sensors = append(sensors, Sensor{common.Point{X: sensorX, Y: sensorY}, common.Point{X: beaconX, Y: beaconY}})
	}
	grid := Grid{
		sensors,
		make(map[int]map[int]bool),
		make(map[int]map[int]bool),
	}

	for _, sensor:= range sensors  {
		grid.add(grid.sensorPoints, sensor.position)
		grid.add(grid.beaconPoints, sensor.closestBeacon)
	}

	return grid
}

func (g Grid) InLine(y int) []common.Range {
	ranges := make([]common.Range, 0)
	for _, sensor := range g.sensors {
		r, inArea := g.InArea(y, sensor)
		if inArea {
			if debug {
				fmt.Printf("Sensor at %v => range %v\n", sensor.position, r)
			}
			ranges = append(ranges, r)
		}
	}
	for noIntersections := false; !noIntersections;  {
		noIntersections = true
		for i := 0; i < len(ranges) && noIntersections; i++ {
			for j := i + 1; j < len(ranges); j++ {
				union := ranges[i].Union(ranges[j])
				if len(union) == 1 {
					if len(ranges) > 2 {
						newRanges := make([]common.Range, 0)
						newRanges = append(newRanges, ranges[:i]...)
						newRanges = append(newRanges, ranges[i + 1:j]...)
						newRanges = append(newRanges, ranges[j + 1:]...)
						ranges = append(newRanges, union...)
					} else {
						ranges = union
					}
					noIntersections = false
					break
				}
			}
		}
	}
	return ranges
}

func (d Day15) SolvePart1(inputFile string, data []string) string {
	grid := d.CreateGrid(inputFile)
	y := 2000000
	if len(data) > 0 {
		v, _ := strconv.Atoi(data[0])
		y = v
	}
	count := 0
	ranges := grid.InLine(y)	
	for _, r := range ranges {
		count += r.Size()
	}
	count = count - grid.NItemsInRow(y)
	return strconv.Itoa(count)

}

func (d Day15) SolvePart2(inputFile string, data []string) string {
	grid := d.CreateGrid(inputFile)
	maxY := 4000000
	if len(data) > 0 {
		v, _ := strconv.Atoi(data[0])
		maxY = v
	}
	var res string = ""
	for i := 0; i <= maxY; i++ {
		if i % 100 == 0 || debug {
			fmt.Printf("%d\n", i)
		}
		ranges := grid.InLine(i)	
		len := len(ranges)
		if len == 2 {
			var lower common.Range
			var upper common.Range
			if ranges[0].Lower() < ranges[1].Lower() {
				lower = ranges[0]
				upper = ranges[1]
			} else {
				lower = ranges[1]
				upper = ranges[0]
			}
			if upper.Lower() - lower.Upper() > 2 {
				//It has to have just one point
				continue
			}
			point := common.Point{X: lower.Upper() + 1, Y: i}
			if !grid.CheckPoint(point) {
				tunningFrequency := (4000000 * point.X) + point.Y
				res = strconv.Itoa(tunningFrequency)
				break
			}
						
		} else if len > 2 {
			fmt.Printf("Range with more than 1 space %d\n", i)
		}
	}	
	return res
}
