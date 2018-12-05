package main

import (
    "./parser"
    "fmt"
)

func main() {
    output := parser.CalculatorInterpret("2 ^ 3 ^ 2")
    fmt.Printf("Output: %f\n", output)
}
