package main

import (
	"fmt"

	"github.com/Marityr/numgo/numgo"
)

func main() {
	A := numgo.NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})

	B := numgo.NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})

	data := A.Sum(B)

	fmt.Println(data.Data)
}
