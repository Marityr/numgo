package matrix

import (
	"errors"
	"log"
)

type Matrixtype interface {
	int64 | float64
}

type Matrix[T Matrixtype] struct {
	Data    [][]T
	capRows int
	capCols int
}

// New create a new matrix
func New[T Matrixtype](data [][]T) Matrix[T] {
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

func Identity[T Matrixtype](n int) Matrix[T] {
	m := make([][]T, n)
	for i := range m {
		m[i] = make([]T, n)
		m[i][i] = 1
	}
	return New(m)
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

// Multiply  matrix * number
func (m Matrix[T]) MultiplyNum(num T) Matrix[T] {
	Subtract, err := m.resultMatrix(m)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < m.capRows; i++ {
		for j := 0; j < m.capCols; j++ {
			Subtract[i][j] = m.Data[i][j] * num
		}
	}

	return New(Subtract)
}

// Multiply matrix power
func (m Matrix[T]) MatrixPower(n int) Matrix[T] {
	result, err := m.resultMatrix(m)
	if err != nil {
		log.Panic(err)
	}

	data := New(result)

	for i := 1; i < n; i++ {
		data = *data.Multiply(data)
	}
	return data
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
		subMatrix := New(make([][]T, size-1))
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

// Rank matrix
func (m Matrix[T]) Rank() T {
	matrix := m.Data
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	rank := 0
	for i := 0; i < cols && rank < rows; i++ {
		pivot := -1
		for j := rank; j < rows; j++ {
			if matrix[j][i] != 0 {
				pivot = j
				break
			}
		}
		if pivot == -1 {
			continue
		}
		if pivot != rank {
			matrix[rank], matrix[pivot] = matrix[pivot], matrix[rank]
		}
		for j := rank + 1; j < rows; j++ {
			coeff := matrix[j][i] / matrix[rank][i]
			for k := i; k < cols; k++ {
				matrix[j][k] -= coeff * matrix[rank][k]
			}
		}
		rank++
	}
	return T(rank)
}

func (m Matrix[T]) InverseMatrix() Matrix[T] {
	n := len(m.Data)
	E := make([][]T, n)
	for i := range E {
		E[i] = make([]T, n)
		E[i][i] = 1
	}
	for k := 0; k < n; k++ {
		Akk := m.Data[k][k]
		for j := 0; j < n; j++ {
			m.Data[k][j] /= Akk
			E[k][j] /= Akk
		}
		for i := 0; i < n; i++ {
			if i == k {
				continue
			}
			Aik := m.Data[i][k]
			for j := 0; j < n; j++ {
				m.Data[i][j] -= Aik * m.Data[k][j]
				E[i][j] -= Aik * E[k][j]
			}
		}
	}
	return New(E)
}

// TransposeMatrix
func (m Matrix[T]) TransposeMatrix() Matrix[T] {
	result, err := m.resultMatrix(m)
	if err != nil {
		log.Panic(err)
	}
	data := New(result)
	transposed := make([][]int, data.capCols)
	for i := range transposed {
		transposed[i] = make([]int, data.capRows)
		for j := range transposed[i] {
			data.Data[i][j] = m.Data[j][i]
		}
	}
	return data
}

// Minor of a matrix
func (m Matrix[T]) Minor(k int) T {
	// Создаем подматрицу из выбранных строк и столбцов
	subMatrix := make([][]T, k)
	for i := range subMatrix {
		subMatrix[i] = make([]T, k)
		for j := range subMatrix[i] {
			subMatrix[i][j] = m.Data[i][j]
		}
	}
	// Вычисляем определитель подматрицы
	data := New(subMatrix)
	return data.Determinant()
}

// Trace of a matrix
func (m Matrix[T]) Trace() T {
	var sum T
	for i := 0; i < len(m.Data); i++ {
		sum += m.Data[i][i]
	}
	return sum
}
