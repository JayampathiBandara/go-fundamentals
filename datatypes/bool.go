package datatypes

import (
	"fmt"
	"go-fundamentals/common"
)

func init() {
	fmt.Println("\nInitializing bools package\n------------------------------")
}

var IsActive bool = true
var IsAdmin bool = false

func PrintBools() {
	fmt.Println("\nBoolean Demo")
	common.PrintLine("_")
	if IsActive {
		fmt.Println("User is active")
	}

	fmt.Println("Is Admin:", IsAdmin)
	common.PrintLine("_")
}
