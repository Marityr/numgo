package examples

import (
	"fmt"

	"github.com/Marityr/numgo/numgo"
)

//

func main() {
	A := [][]int{[]int{1, 2}, []int{3, 4}}
	B := [][]int{[]int{5, 6}, []int{7, 8}}

	fmt.Println(numgo.Matrix.Sum(A, B))
}
