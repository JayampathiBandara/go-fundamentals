package datatypes

import (
	"errors"
	"fmt"
)

var SampleError error = errors.New("this is a sample error")

func init() {
	println("\nInitializing error package\n------------------------------")
}

func PrintError(i int) error {
	fmt.Println("\nError Demo\n------------------------------")
	if i == 0 {
		return SampleError
	}
	return nil
}
