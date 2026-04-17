package aggregatedatatype

import (
	"fmt"
	"go-fundamentals/common"
)

func ArrayDeclaration() {
	// Arrays in Go are fixed-size collections of elements of the same type.
	// The size of the array is determined at compile time and cannot be changed.
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println(numbers) // [10 20 0 0 0]

	// ... is used to let the compiler determine the size of the array based on the number of elements provided.
	scores := [...]int{90, 85, 70, 100}
	fmt.Println(scores)

	for i := 0; i < len(scores); i++ {
		fmt.Println(i, scores[i])
	}

	for index, value := range scores {
		fmt.Println(index, value)
	}
	common.PrintLine("=")
}
