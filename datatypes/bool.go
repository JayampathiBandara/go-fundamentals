package datatypes

import "fmt"

func init() {
	fmt.Println("\nInitializing bools package\n------------------------------")
}

var IsActive bool = true
var IsAdmin bool = false

func PrintBools() {
	fmt.Println("\nBoolean Demo\n------------------------------")
	if IsActive {
		fmt.Println("User is active")
	}

	fmt.Println("Is Admin:", IsAdmin)
}
