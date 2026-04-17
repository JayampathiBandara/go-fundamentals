package datatypes

import "fmt"

func init() {
	fmt.Println("\nInitializing strings package\n------------------------------")
}

var FirstName string = "John"
var LastName string = "Smith"

func PrintStrings() {
	fmt.Println("\nString Demo\n------------------------------")
	fmt.Println("Full Name", FirstName, LastName)
	fmt.Println(`Address: 
	122A Sengkang Eastway,
	Singapore.`)
	Age = 45

	fmt.Println("Age strings", Age)
}
