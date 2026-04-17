package pointers

import "go-fundamentals/common"

var originalValue int = 10

var pointerToValue = &originalValue

var pointerToPointer = &pointerToValue

var valueReferencedFromPointer = *pointerToValue
var valueReferencedOfOriginal = originalValue

func Printvariables(status string) {
	println("\nPointers Demo :", status)
	common.PrintLine("-")
	println("Original Value: ", originalValue)
	println("Pointer to Value: ", pointerToValue)
	println("Pointer to Pointer: ", pointerToPointer)
	println("Pointer to Value (Use * to Dereference Value): ", *pointerToValue)
	println("Pointer to Pointer: (Use * to Dereference Value)", *pointerToPointer)
	println("Pointer to Pointer: (Use ** to Dereference Value)", **pointerToPointer)
	println("Value Referenced from Pointer: ", valueReferencedFromPointer)
	println("Value Referenced of Original: ", valueReferencedOfOriginal)
}

func UpdateOriginalValue(newValue int) {
	originalValue = newValue
}
