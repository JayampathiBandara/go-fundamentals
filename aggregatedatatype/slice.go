package aggregatedatatype

import (
	"fmt"
	"go-fundamentals/common"
	"slices"
)

func SliceDeclaration() {
	common.PrintLine("=")
	fmt.Println(`Slices in Go are dynamic, flexible views into the elements of an array. 
	They are more commonly used than arrays because they provide more functionality and are easier to work with.`)
	var s []int
	fmt.Println(s)
	s = []int{1, 2, 3}
	fmt.Println(s[1])

	s[1] = 99
	fmt.Println(s)
	s = append(s, 5, 10, 15)
	fmt.Println("After call Append", s)

	s = slices.Delete(s, 1, 3)
	fmt.Println("After call Delete", s)
	common.PrintLine("=")
}
