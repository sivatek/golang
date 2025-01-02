// The dot product of two vectors is calculated by this formula:
// DOT PRODUCT = A*B
//             = Ax * Bx + Ay * By + Az * Bz
// where A = (1, 2, 3) and B = (4, 5, 6)
// then  A.B = (1.4 + 2.5 + 3.6) = 32

package main

import "fmt"

func dotProduct(vec1 []int, vec2 []int) int {
	if len(vec1) != len(vec2) {
		panic("Vectors must have the same length")
	}
	dotProduct := 0
	for i := 0; i < len(vec1); i++ {
		dotProduct += vec1[i] * vec2[i]
	}
	return dotProduct
}

func main() {
	vec1 := []int{1, 2, 3}
	vec2 := []int{4, 5, 6}
	dotProduct := dotProduct(vec1, vec2)
	fmt.Println("The dot product is: ", dotProduct)
}
