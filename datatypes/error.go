package datatypes

import (
	"errors"
	"fmt"
	"go-fundamentals/common"
)

var SampleError error = errors.New("this is a sample error")

func init() {
	println("\nInitializing error package")
	common.PrintLine("_")
}

func PrintError(i int) error {
	fmt.Println("\nError Demo")
	common.PrintLine("_")
	if i == 0 {
		return SampleError
	}
	return nil
}
