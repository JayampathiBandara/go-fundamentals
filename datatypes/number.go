package datatypes

import (
	"fmt"
	"go-fundamentals/common"
	"math"
)

var Age int = 40

const Height float32 = 1.44 // m
var Weight int = 100        // Kg

func init() {
	println("\nInitializing numbers package")
	common.PrintLine("_")
}

func PrintNumbers() {
	fmt.Println("\nNumbers Demo")
	common.PrintLine("_")
	println("Age (numbers): ", Age)
	println("Height (numbers): ", Height)
	println("Weight (numbers): ", Weight)
	println("BMI : ", calculateBMI(Height, Weight))
}

func calculateBMI(height float32, weight int) float32 {
	h := float64(height)
	w := float64(weight)

	bmi := w / math.Pow(h, 2)

	return float32(bmi)
}
