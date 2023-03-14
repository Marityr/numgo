package matrix

import (
	"fmt"
	"testing"
)

var (
	A = New([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})

	B = New([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})

	AError = New([][]float64{
		{1, 2, 3, 4},
	})
	BError = New([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7},
	})
)

func TestSumMatrix(t *testing.T) {
	t.Run("sum two matrices", func(t *testing.T) {
		got := A.Sum(B)
		want := New([][]float64{
			{2, 4, 6, 8},
			{10, 12, 14, 16},
		})
		assertMatrixEqual(t, *got, want)
	})

	t.Run("panic error Sum", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("want panic, got nil")
			}
		}()
		A.Sum(AError)
	})
}

func TestSubtractMatrix(t *testing.T) {
	t.Run("subtract two matrices", func(t *testing.T) {
		got := A.Subtract(B)
		want := New([][]float64{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		})
		assertMatrixEqual(t, *got, want)
	})

	t.Run("panic error Subtract", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("want panic, got nil")
			}
		}()
		A.Subtract(AError)
	})
}

func TestMultiplyMatrix(t *testing.T) {
	t.Run("multiply two matrices", func(t *testing.T) {
		got := A.Multiply(B)
		want := New([][]float64{
			{1, 4, 9, 16},
			{25, 36, 49, 64},
		})
		assertMatrixEqual(t, *got, want)
	})

	t.Run("panic error Multiply", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("want panic, got nil")
			}
		}()
		A.Multiply(AError)
	})
}

func TestResultMatrix(t *testing.T) {
	t.Run("result nil error", func(t *testing.T) {
		_, want := A.resultMatrix(B)

		if want != nil {
			t.Errorf("want nil, got %v", want)
		}
	})

	t.Run("result err error", func(t *testing.T) {
		_, LenError := A.resultMatrix(AError)
		_, ColError := A.resultMatrix(BError)

		assertLenError(t, LenError)
		assertColError(t, ColError)
	})
}

func TestNewMatrix(t *testing.T) {
	t.Run("new matrix", func(t *testing.T) {
		got := New([][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		})
		want := New([][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		})
		assertMatrixEqual(t, got, want)
	})
}

func TestDeterminant(t *testing.T) {
	t.Run("determinant 3x3", func(t *testing.T) {
		got := New([][]int64{
			{3, 3, 3},
			{3, 3, 3},
			{3, 3, 3},
		}).Determinant()

		want := int64(0)
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("determinant 3x1", func(t *testing.T) {
		got := New([][]int64{
			{3, 3, 3},
		}).Determinant()

		want := int64(3)
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("determinant 0", func(t *testing.T) {
		got := New([][]float64{}).Determinant()
		want := float64(0)

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestRank(t *testing.T) {
	t.Run("rank 3x3", func(t *testing.T) {
		got := New([][]int64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}).Rank()

		want := int64(2)

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}

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

func assertMatrixEqual[T Matrixtype](t *testing.T, got Matrix[T], want Matrix[T]) {
	t.Helper()

	for i := 0; i < got.capRows; i++ {
		for j := 0; j < got.capCols; j++ {
			if got.Data[i][j] != want.Data[i][j] {
				t.Errorf("got %v, want %v", got.Data[i][j], want.Data[i][j])
			}
		}
	}
}

func assertMatrixNil(t *testing.T, got *Matrix[float64]) {
	t.Helper()

	if got != nil {
		t.Errorf("want nil, got %v", got)
	}
}

func ExampleMatrix() {
	a := Matrix[float64]{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	b := Matrix[float64]{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}
	repead := a.Sum(b)
	fmt.Println(repead)
	// Output: &{[[2 4 6 8] [10 12 14 16]] 2 4}
}
