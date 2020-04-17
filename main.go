//
// High-level modules should not depend on low-level modules. Both should depend on abstractions.
// Abstractions should not depend on details. Details should depend on abstractions.
// Objects (or software entities) should be open for extension, but closed for modification.

package main

import "fmt"

// Calculator calculate :)
type Calculator struct {
	operation Operation
}

// Solve calculates the operation between two numbers
func (calculator Calculator) Solve(a float32, b float32) float32 {
	return calculator.operation.Calculate(a, b)
}

// Operation defines calculation operation behaviour
type Operation interface {
	Calculate(float32, float32) float32
}

// NewCalculator creates new calculator
func NewCalculator(op Operation) Calculator {
	return Calculator{operation: op}
}

// AddOperation adds two numbers
type AddOperation struct {
}

// Calculate adds two numbers
func (addOperation AddOperation) Calculate(a float32, b float32) float32 {
	return a + b
}

// SubtractOperation subtract two numbers
type SubtractOperation struct {
}

// Calculate subtract two numbers
func (subtractOperation SubtractOperation) Calculate(a float32, b float32) float32 {
	return a - b
}

func main() {

	calculator := NewCalculator(AddOperation{})

	fmt.Println(calculator.Solve(2, 2))

	calculator = NewCalculator(SubtractOperation{})

	fmt.Println(calculator.Solve(2, 2))

}
