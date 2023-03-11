package numgo

import (
	"errors"
)

type Matrix struct{}

// Summation of two matrices
func (m *Matrix) Sum(A, B [][]int) (sum [][]int, err error) {
	if len(A) != len(B) {
		return nil, errors.New("matrix.Sum: len(a) != len(b)")
	}

	rows := len(A)
	cols := len(A[0])

	Sum := make([][]int, rows)
	for i := range Sum {
		Sum[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			Sum[i][j] = A[i][j] + B[i][j]
		}
	}

	return Sum, nil
}
