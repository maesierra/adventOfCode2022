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
				columns++
			}
			data = append(data, float64(ch-'0'))
		}
		columnsCalculated = true
		rows++
	}
	return mat.NewDense(rows, columns, data)
}

func AddColumnToTheLeft(m *mat.Dense) *mat.Dense {
	newData := []float64{}
	for r := 0; r < m.RawMatrix().Rows; r++ {
		newData = append(newData, 0)
		for c := 0; c < m.RawMatrix().Cols; c++ {
			newData = append(newData, m.At(r, c))
		}
	}
	return mat.NewDense(m.RawMatrix().Rows, m.RawMatrix().Cols+1, newData)
}

func AddColumnToTheRight(m *mat.Dense) *mat.Dense {
	return mat.DenseCopyOf(m.Grow(0, 1))
}

func AddRowToTheBottom(m *mat.Dense) *mat.Dense {
	return mat.DenseCopyOf(m.Grow(1, 0))
}

func AddRowToTheTop(m *mat.Dense) *mat.Dense {
	newRow := make([]float64, m.RawMatrix().Cols)
	return mat.NewDense(m.RawMatrix().Rows+1, m.RawMatrix().Cols, append(newRow, m.RawMatrix().Data...))
}
