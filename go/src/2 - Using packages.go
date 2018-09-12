// Naming this your main package, if this isn't named main, code won't run
package main

// Importing the relevant packages

import (
  "fmt"                 // Same package we previously saw
  m "./mymath"            // m is an alias for the my_math package we previosuly defined
  "math"                // In built math package
)

// defining main function
func main() {
  var a int = 3             // Explicit type decalaration
  b := 5                    // Implicit type declaration

  fmt.Print("Enter 2 numbers: ")
  fmt.Scanf("%d %d", &a, &b)      // Same as C

  sum := m.Sum(a, b)
  diff := m.Diff(a, b)
  mul := m.Mul(a, b)
  div := m.Div(float64(a), float64(b))
  sqrt := math.Sqrt(float64(a))
  fmt.Printf("Sum: %d\nDifference: %d\nProduct: %d\nDivision: %f\nSqrt(%d): %f", sum, diff, mul, div, a, sqrt)
}
