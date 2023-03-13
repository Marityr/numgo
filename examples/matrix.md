
```go
    A := numgo.Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	B := numgo.Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	sum := A.Sum(B.Data)

	fmt.Println(sum)

	C := numgo.Matrix{
		Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}

	sum = A.Sum(B.Sum(C.Data))

	fmt.Println(sum)
```