package numgo

import (
	"errors"
	"log"
)

type matrixtype interface {
	int64 | float64
}

type Matrix[T matrixtype] struct {
	Data    [][]T
	capRows int
	capCols int
}

func NewMatrix[T matrixtype](data [][]T) Matrix[T] {
	if len(data) == 0 {
		return Matrix[T]{
			Data:    [][]T{},
			capRows: 0,
			capCols: 0,
		}
	}
	return Matrix[T]{
		Data:    data,
		capRows: len(data),
		capCols: len(data[0]),
	}
}

func (m Matrix[T]) resultMatrix(B Matrix[T]) ([][]T, error) {
	if len(m.Data) != len(B.Data) {
		return nil, errors.New("matrix dimensions are not equal")
	}

	for i := range m.Data {
		if len(m.Data[i]) != len(B.Data[i]) {
			return nil, errors.New("matrix dimensions are not equal")
		}
	}

	m.capRows = len(m.Data)
	m.capCols = len(B.Data[0])

	result := make([][]T, m.capRows)
	for i := range result {
		result[i] = make([]T, m.capCols)
	}

	return result, nil
}

// Summation of two matrices
//
// A + B
//
// A.Sum(B)
func (m Matrix[T]) Sum(B Matrix[T]) *Matrix[T] {
	Sum, err := m.resultMatrix(B)
	if err != nil {
		log.Panic(err)
		return nil
	}

	for i := 0; i < m.capRows; i++ {
		for j := 0; j < m.capCols; j++ {
			Sum[i][j] = m.Data[i][j] + B.Data[i][j]
		}
	}

	return &Matrix[T]{Sum, m.capRows, m.capCols}
}

// Summation of two matrices
// A - B
// A.Subtract(B)
func (m Matrix[T]) Subtract(B Matrix[T]) *Matrix[T] {
	Subtract, err := m.resultMatrix(B)
	if err != nil {
		log.Panic(err)
		return nil
	}

	for i := 0; i < m.capRows; i++ {
		for j := 0; j < m.capCols; j++ {
			Subtract[i][j] = m.Data[i][j] - B.Data[i][j]
		}
	}

	return &Matrix[T]{Subtract, m.capRows, m.capCols}
}

// Multiplication of two matrices
// A * B
// A.Multiply(B)
func (m Matrix[T]) Multiply(B Matrix[T]) *Matrix[T] {
	Multiply, err := m.resultMatrix(B)
	if err != nil {
		log.Panic(err)
		return nil
	}

	for i := 0; i < m.capRows; i++ {
		for j := 0; j < m.capCols; j++ {
			Multiply[i][j] = m.Data[i][j] * B.Data[i][j]
		}
	}

	return &Matrix[T]{Multiply, m.capRows, m.capCols}
}

// Determinant of a matrix
func (m Matrix[T]) Determinant() T {
	matrix := m.Data
	size := len(matrix)
	if size < 1 {
		return 0
	} else if size == 1 {
		return matrix[0][0]
	} else if size == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}
	var det T
	for i := 0; i < size; i++ {
		subMatrix := NewMatrix(make([][]T, size-1))
		for j := range subMatrix.Data {
			subMatrix.Data[j] = make([]T, size-1)
		}
		for j := 1; j < size; j++ {
			for k := 0; k < size; k++ {
				if k < i {
					subMatrix.Data[j-1][k] = matrix[j][k]
				} else if k > i {
					subMatrix.Data[j-1][k-1] = matrix[j][k]
				}
			}
		}
		sign := 1.0
		if i%2 != 0 {
			sign = -1.0
		}
		det += T(sign) * matrix[0][i] * subMatrix.Determinant()
	}
	return T(det)
}
