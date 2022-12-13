package day12

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
	"maesierra.net/advent-of-code/2022/common"
)

type Day12 struct {
}

type HeightMap struct {
	heightMap dijkstra.Graph
	start string
	end string
	lowerPoints [] string
}

func (h HeightMap) Path(start, end string) (path []string, cost int, err error){
	return h.heightMap.Path(start, end)
}

func heightDiff(ch1, ch2 rune) int {
	if ch1 == 'E' {
		ch1 = 'z'
	} else if ch2 == 'E' {
		ch2 = 'z'
	}
	if ch1 == 'S' {
		ch1 = 'a'
	} else if ch2 == 'S' {
		ch2 = 'a'
	}
	diff := int(ch1) - int(ch2)
	return diff
}

func key(row, col int) string {
	return fmt.Sprintf("%d-%d", row, col)
}

func (d Day12) CreateHeightMap(inputFile string) HeightMap {
	input := strings.Split(common.ReadFile(inputFile), "\n")
	rows := len(input)
	cols := len(input[0])
	heightMap := dijkstra.Graph{}
	start := ""
	end := ""
	lowerPoints := []string{}
	for row, line := range input {
		for col, ch := range line {
			if ch == 'E' {
				end = key(row, col)
			}
			if ch == 'S' {
				start = key(row, col)
			}
			if ch == 'S' || ch == 'a' {
				lowerPoints = append(lowerPoints, key(row, col))
			}
			neighbours := map[string]int{}
			//up
			if row > 0 {
				diff := heightDiff(rune(input[row-1][col]), ch)
				if diff < 2 {
					neighbours[key(row -1, col)] = 1
				}
			}
			//left
			if col < cols -1  {
				diff := heightDiff(rune(input[row][col+1]), ch)
				if diff < 2 {
					neighbours[key(row, col + 1)] = 1
				}
			}
			//down  
			if row < rows -1 {
				diff := heightDiff(rune(input[row+1][col]), ch)
				if (diff < 2) {
					neighbours[key(row + 1, col)] = 1
				}
			}
			//right
			if col > 0 {
				diff := heightDiff(rune(input[row][col-1]), ch)
				if (diff < 2) {
					neighbours[key(row, col - 1)] = 1
				}
			}
			heightMap[key(row, col)] = neighbours
		}
	}
	return HeightMap{heightMap: heightMap, start: start, end: end, lowerPoints: lowerPoints}
}

func (d Day12) SolvePart1(inputFile string) string {
	
	heightMap := d.CreateHeightMap(inputFile)
	path, cost, _ := heightMap.Path(heightMap.start, heightMap.end)
	fmt.Printf("path: %v, cost: %v, start: %v end %v\n", path, cost, heightMap.start, heightMap.end)
	return strconv.Itoa(cost)

}

func (d Day12) SolvePart2(inputFile string) string {
	heightMap := d.CreateHeightMap(inputFile)
	minCost := math.MaxInt
	for i, pos := range heightMap.lowerPoints {
		fmt.Printf("Solving %v %d/%d\n", pos, i + 1, len(heightMap.lowerPoints))
		path, cost, _ := heightMap.Path(pos, heightMap.end)
		if len(path) > 1 {
			minCost = common.IntMin(minCost, cost)
		}
	}
	return strconv.Itoa(minCost)
}
