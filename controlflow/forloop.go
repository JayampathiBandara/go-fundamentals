package controlflow

import (
	"fmt"
	"go-fundamentals/common"
	"strconv"
)

func ForLoop() {
	// For loop is the only loop construct in Go. It can be used to iterate over a range of values,
	// or to execute a block of code a specific number of times.
	for i := 0; i < 5; i++ {
		println(i)
	}
	common.PrintLine("=")
	// You can also use a for loop without the initialization and post statements,
	// similar to a while loop in other languages.
	j := 0
	for j < 5 {
		println(j)
		j++
	}
	common.PrintLine("=")
	// You can also use a for loop with a range clause to iterate over elements in a collection,
	// such as an array or slice.
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		println(index, value)
	}
	common.PrintLine("=")
	type score struct {
		name  string
		score int
	}

	scores := []score{}

	for {
		fmt.Println("Enter a student name and score (or type 'exit' to stop):")

		var name, rawScore string
		fmt.Scanln(&name, &rawScore)

		// 🔹 Break condition
		if name == "exit" {
			break
		}

		s, err := strconv.Atoi(rawScore)
		if err != nil {
			fmt.Println("Invalid score, try again")
			continue
		}

		scores = append(scores, score{name: name, score: s})
	}
	fmt.Println("Scores entered:")
	for _, s := range scores {
		fmt.Printf("%s: %d\n", s.name, s.score)
	}

	common.PrintLine("*")
}
