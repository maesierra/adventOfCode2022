package day8

import (
	"fmt"
	"strconv"

	"gonum.org/v1/gonum/mat"
	"maesierra.net/advent-of-code/2022/common"
	"maesierra.net/advent-of-code/2022/common/intmatrix"
)

type Day8 struct {
}

func (d Day8) SolvePart1(inputFile string) string {

	treeMap := intmatrix.ReadFileIntoMatrix(inputFile)
	
	rows, columns := treeMap.RawMatrix().Rows, treeMap.RawMatrix().Cols
	visibleTrees := (rows * 2) + (columns * 2) - 4 // The trees on the border are always visible


	//We iterate excluding the borders
	for r := 1; r < rows - 1; r++ {
		for c := 1; c < columns - 1; c++ {			
			current := treeMap.At(r, c)
						
			if d.IsHigherThanSlice(current, d.SliceRow(treeMap, r, 0, c)) {
				visibleTrees++
				continue
			}
			if d.IsHigherThanSlice(current, d.SliceRow(treeMap, r, c + 1, columns)) {
				visibleTrees++
				continue
			}
			
			if d.IsHigherThanSlice(current, d.SliceCol(treeMap, c, 0, r)) {
				visibleTrees++
				continue
			}
			if d.IsHigherThanSlice(current, d.SliceCol(treeMap, c, r + 1, rows)) {
				visibleTrees++
				continue
			}
		}		
	}

	return strconv.Itoa(visibleTrees)

}

func (d Day8) SolvePart2(inputFile string) string {
	treeMap := intmatrix.ReadFileIntoMatrix(inputFile)
	
	rows, columns := treeMap.RawMatrix().Rows, treeMap.RawMatrix().Cols

	maxScenicScore := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {			
			current := treeMap.At(r, c)
			scenicScore := 1

			up    :=  d.SliceCol(treeMap, c, 0, r)
			upScore := d.CalculateVisibleTrees(common.Reverse(up), current)
			scenicScore *= upScore

			right  :=  d.SliceRow(treeMap, r, c + 1, columns)
			rightScore := d.CalculateVisibleTrees(right, current)
			scenicScore *=  rightScore

			bottom :=  d.SliceCol(treeMap, c, r + 1, rows)
			bottomScore := d.CalculateVisibleTrees(bottom, current)
			scenicScore *=  bottomScore

			left   :=  d.SliceRow(treeMap, r, 0, c)
			leftScore := d.CalculateVisibleTrees(common.Reverse(left), current)
			scenicScore *=  leftScore
			if (scenicScore > maxScenicScore) {
				fmt.Printf("New max found at %v %v => %v\n", r, c, scenicScore)
				maxScenicScore = scenicScore
			}
			
		}		
	}

	return strconv.Itoa(maxScenicScore)
}

func (d Day8) CalculateVisibleTrees(slice []float64, current float64) int {
	c := 0
	for _, v := range slice {
		if v >= current {
			c++;
			break
		}
		c++
	}
	return c
}

func (d Day8) IsHigherThanSlice(value float64, slice []float64) bool {
	res := true								
	for _, v:= range slice {
		if v >= value {
			return false
		}
	}
	return res

}

func (d Day8) SliceVector(vector mat.Vector, c1, c2 int) []float64 {
	slice := []float64{}
	for idx := common.IntMax(c1, 0); idx < common.IntMin(c2, vector.Len()); idx++ {
		slice = append(slice, vector.AtVec(idx))
	}
	return slice
}

func (d Day8) SliceRow(matrix * mat.Dense, r, c1, c2 int) []float64 {
	return d.SliceVector(matrix.RowView(r), c1, c2)
}
func (d Day8) SliceCol(matrix * mat.Dense, c, r1, r2 int) []float64 {
	return d.SliceVector(matrix.ColView(c), r1, r2)
}

