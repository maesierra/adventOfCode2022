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

func AddColumnToTheLeft(n int, m *mat.Dense) *mat.Dense {
	newData := []float64{}
	filled := []float64{}
	for i := 0; i < n; i++ {
		filled = append(filled, 0)
	}
	for r := 0; r < m.RawMatrix().Rows; r++ {
		newData = append(newData, filled...)
		for c := 0; c < m.RawMatrix().Cols; c++ {
			newData = append(newData, m.At(r, c))
		}
	}
	return mat.NewDense(m.RawMatrix().Rows, m.RawMatrix().Cols + n, newData)
}

func AddColumnToTheRight(n int, m *mat.Dense) *mat.Dense {
	return mat.DenseCopyOf(m.Grow(0, n))
}

func AddRowToTheBottom(n int, m *mat.Dense) *mat.Dense {
	return mat.DenseCopyOf(m.Grow(n, 0))
}

func AddRowToTheTop(n int, m *mat.Dense) *mat.Dense {
	newRow := make([]float64, m.RawMatrix().Cols)
	data := m.RawMatrix().Data
	for i := 0; i < n; i++ {
		data = append(newRow, data...)
	}
	return mat.NewDense(m.RawMatrix().Rows + n, m.RawMatrix().Cols, data)
}
