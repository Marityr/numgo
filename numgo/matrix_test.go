package numgo

import (
	"fmt"
	"testing"
)

var (
	A = Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	B = Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	AError = Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
		},
	}
	BError = Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7},
		},
	}
)

func TestSumMatrix(t *testing.T) {
	t.Run("sum two matrices", func(t *testing.T) {
		got := A.Sum(B.Data)
		want := Matrix{
			Data: [][]float64{
				{2, 4, 6, 8},
				{10, 12, 14, 16},
			},
		}
		assertMatrixEqual(t, got, want)
	})

	t.Run("panic error Sum", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("want panic, got nil")
			}
		}()
		A.Sum(AError.Data)
	})
}

func TestSubtractMatrix(t *testing.T) {
	t.Run("subtract two matrices", func(t *testing.T) {
		got := A.Subtract(B.Data)
		want := Matrix{
			Data: [][]float64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		}
		assertMatrixEqual(t, got, want)
	})

	t.Run("panic error Subtract", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("want panic, got nil")
			}
		}()
		A.Subtract(AError.Data)
	})
}

func TestMultiplyMatrix(t *testing.T) {
	t.Run("multiply two matrices", func(t *testing.T) {
		got := A.Multiply(B.Data)
		want := Matrix{
			Data: [][]float64{
				{1, 4, 9, 16},
				{25, 36, 49, 64},
			},
		}
		assertMatrixEqual(t, got, want)
	})

	t.Run("panic error Multiply", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("want panic, got nil")
			}
		}()
		A.Multiply(AError.Data)
	})
}

func TestResultMatrix(t *testing.T) {
	t.Run("result nil error", func(t *testing.T) {
		_, want := A.resultMatrix(B.Data)

		if want != nil {
			t.Errorf("want nil, got %v", want)
		}
	})

	t.Run("result err error", func(t *testing.T) {
		_, LenError := A.resultMatrix(AError.Data)
		_, ColError := A.resultMatrix(BError.Data)

		assertLenError(t, LenError)
		assertColError(t, ColError)
	})
}

func assertLenError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Error("want error, got nil")
	}
}

func assertColError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Error("want error, got nil")
	}
}

func assertMatrixEqual(t *testing.T, got *Matrix, want Matrix) {
	t.Helper()

	for i := 0; i < got.capRows; i++ {
		for j := 0; j < got.capCols; j++ {
			if got.Data[i][j] != want.Data[i][j] {
				t.Errorf("got %v, want %v", got.Data[i][j], want.Data[i][j])
			}
		}
	}
}

func assertMatrixNil(t *testing.T, got *Matrix) {
	t.Helper()

	if got != nil {
		t.Errorf("want nil, got %v", got)
	}
}

func ExampleMatrix() {
	a := Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	b := Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}
	repead := a.Sum(b.Data)
	fmt.Println(repead)
	// Output: &{[[2 4 6 8] [10 12 14 16]] 2 4}
}
