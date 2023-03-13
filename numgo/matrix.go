package numgo

import (
	"errors"
	"log"
)

type DataMatrix interface {
	Sum(B [][]float64) *Matrix
	Subtract(B [][]float64) *Matrix
	Multiply(B [][]float64) *Matrix
}

type Matrix struct {
	Data    [][]float64
	capRows int
	capCols int
}

func (m *Matrix) resultMatrix(B [][]float64) ([][]float64, error) {
	if len(m.Data) != len(B) {
		return nil, errors.New("matrix dimensions are not equal")
	}

	for i := range m.Data {
		if len(m.Data[i]) != len(B[i]) {
			return nil, errors.New("matrix dimensions are not equal")
		}
	}

	m.capRows = len(m.Data)
	m.capCols = len(B[0])

	result := make([][]float64, m.capRows)
	for i := range result {
		result[i] = make([]float64, m.capCols)
	}

	return result, nil
}

// Summation of two matrices
//
// A + B
//
// A.Sum(B.Data)
func (m *Matrix) Sum(B [][]float64) *Matrix {
	Sum, err := m.resultMatrix(B)
	if err != nil {
		log.Panic(err)
		return nil
	}

	for i := 0; i < m.capRows; i++ {
		for j := 0; j < m.capCols; j++ {
			Sum[i][j] = m.Data[i][j] + B[i][j]
		}
	}

	return &Matrix{Sum, m.capRows, m.capCols}
}

// Summation of two matrices
// A - B
// A.Subtract(B.Data)
func (m *Matrix) Subtract(B [][]float64) *Matrix {
	Subtract, err := m.resultMatrix(B)
	if err != nil {
		log.Panic(err)
		return nil
	}

	for i := 0; i < m.capRows; i++ {
		for j := 0; j < m.capCols; j++ {
			Subtract[i][j] = m.Data[i][j] - B[i][j]
		}
	}

	return &Matrix{Subtract, m.capRows, m.capCols}
}

// Multiplication of two matrices
// A * B
// A.Multiply(B.Data)
func (m *Matrix) Multiply(B [][]float64) *Matrix {
	Multiply, err := m.resultMatrix(B)
	if err != nil {
		log.Panic(err)
		return nil
	}

	for i := 0; i < m.capRows; i++ {
		for j := 0; j < m.capCols; j++ {
			Multiply[i][j] = m.Data[i][j] * B[i][j]
		}
	}

	return &Matrix{Multiply, m.capRows, m.capCols}
}
