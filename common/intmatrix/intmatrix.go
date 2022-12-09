package intmatrix

import (
	"gonum.org/v1/gonum/mat"
	"maesierra.net/advent-of-code/2022/common"
)

func ReadFileIntoMatrix(path string) *mat.Dense {
	var data []float64
	rows := 0
	columns := 0
	columnsCalculated := false
	for _, l := range common.ReadFileIntoLines(path) {
		for _, ch := range l {
			if !columnsCalculated {
				columns++;
			}
			data = append(data, float64(ch-'0'))
		}
		columnsCalculated = true
		rows++
	}
	return mat.NewDense(rows, columns, data)
}
