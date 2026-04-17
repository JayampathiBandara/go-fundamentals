package main

import (
	"fmt"
	"go-fundamentals/aggregatedatatype"
	"go-fundamentals/common"
	"go-fundamentals/pointers"
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

	/*
		datatypes.PrintStrings()
		datatypes.PrintNumbers()
		fmt.Println("Error Message :", datatypes.PrintError(0))
		datatypes.PrintBools()
	*/
	common.PrintLine("=")
	println("values Pass by reference is chaged ufter each update")
	println("But Value pass as a value doesn't change as it is bind at compile time")
	pointers.Printvariables("BEFORE update")
	pointers.UpdateOriginalValue(1000)
	pointers.Printvariables("AFTER update")
	pointers.UpdateOriginalValue(9000)
	pointers.Printvariables("AFTER Second time update")
	common.PrintLine("=")

	aggregatedatatype.ArrayDeclaration()
	aggregatedatatype.SliceDeclaration()
	aggregatedatatype.MapDeclaration()
	aggregatedatatype.VarStructDeclaration()
	aggregatedatatype.TypeStructDeclaration()

}
