package main

import "fmt"

func main() {
  fmt.Println("=== Panic Basics ===")

  // Normal execution before panic
  fmt.Println("Starting program")

  // Call function that panics
  safeDivide(10, 2)
  safeDivide(10, 0) // This will panic

  // This line never runs
  fmt.Println("This never prints")
}

func safeDivide(a, b int) {
  if b == 0 {
    panic("cannot divide by zero")
  }
  result := a / b
  fmt.Printf("%d / %d = %d\n", a, b, result)
}
