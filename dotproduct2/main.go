package main

import "fmt"

func dotProduct(vector1 []int, vector2 []int) int {
	if len(vector1) != len(vector2) {
		panic("Vectors must have the same length")
	}
	dotProduct := 0
	for i, v := range vector1 {
		dotProduct += v * vector2[i]
	}
	return dotProduct
}

func main() {
	vector1 := []int{1, 2, 3}
	vector2 := []int{4, 5, 6}
	dotProduct := dotProduct(vector1, vector2)
	fmt.Println("The Dot Product Is: ", dotProduct)
}
