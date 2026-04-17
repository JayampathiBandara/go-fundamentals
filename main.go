package main

import (
	"fmt"
	"go-fundamentals/datatypes"
)

func init() {
	println(`
	Each package can have an init function which will be called before the main function is called. 
	You can have multiple init functions in a package and they will be called in alphebetical order of file name
	and the order they are defined.`)
	println("\nFinally call main file init function")
	println("Initializing main package\n------------------------------")
}

func main() {
	// call built in print
	println("Happy new year 2026")

	// use packages to call functions
	fmt.Println("Happy New Year - 2026")

	datatypes.PrintStrings()
	datatypes.PrintNumbers()
	fmt.Println("Error Message :", datatypes.PrintError(0))
	datatypes.PrintBools()
}
