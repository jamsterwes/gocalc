package main

import (
    "./interpreter"
    "fmt"
)

func main() {
    output := interpreter.CalculatorInterpret("2 ^ 3 ^ 2")
    fmt.Printf("Output: %f\n", output)
}
