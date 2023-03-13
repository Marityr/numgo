package main

import (
	"fmt"

	"github.com/Marityr/numgo/numgo"
)

func main() {
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

	data := A.Multiply(B.Data)

	fmt.Println(data)
}
