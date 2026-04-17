package aggregatedatatype

import "go-fundamentals/common"

var User struct {
	name string
	age  int
}

type UserStruct struct {
	Name string
	Age  int
}

func VarStructDeclaration() {
	// Following is Not possible
	// u2 := User{}
	// Because User is a variable, not a type.
	// can be define only once for the applicatinn, and it is a global variable that can be accessed and modified
	// from anywhere in the code.

	User.name = "John Doe"
	User.age = 30
	println("User Name:", User.name)
	println("User Age:", User.age)
	common.PrintLine("=")
}

func TypeStructDeclaration() {
	user := UserStruct{Name: "Jane Doe", Age: 25}
	println("User Name:", user.Name)
	println("User Age:", user.Age)

	// You can also declare a struct without initializing it, and then set the fields later.
	var anotherUser UserStruct
	anotherUser.Name = "Alice"
	anotherUser.Age = 28
	println("Another User Name:", anotherUser.Name)
	println("Another User Age:", anotherUser.Age)

	common.PrintLine("=")
}
