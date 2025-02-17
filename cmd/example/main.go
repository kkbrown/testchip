package main

import (
	"fmt"
	"github.com/kkbrown/testchip/pkg/calculator"
	"github.com/kkbrown/testchip/pkg/validator"
)

func main() {
	calc := calculator.NewCalculator(2)
	valid := validator.NewValidator(100)

	result := calc.Add(5, 3)
	isValid := valid.IsValid(result)

	fmt.Printf("Result: %v, IsValid: %v\n", result, isValid)
}
