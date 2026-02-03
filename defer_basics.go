package main

import "fmt"

func main() {
  fmt.Println("=== Defer Basics ===")

  // Defer runs when function returns
  fmt.Println("Start")
  defer fmt.Println("Deferred: runs last")
  fmt.Println("Middle")
  fmt.Println("End")

  fmt.Println("\n=== Multiple Defers (LIFO) ===")
  multipleDefers()

  fmt.Println("\n=== Defer with Arguments ===")
  deferArguments()
}

func multipleDefers() {
  // Defers execute in LIFO order (last in, first out)
  defer fmt.Println("First defer (runs third)")
  defer fmt.Println("Second defer (runs second)")
  defer fmt.Println("Third defer (runs first)")
  fmt.Println("Function body")
}

func deferArguments() {
  // Arguments are evaluated when defer is called, not when it runs
  x := 10
  defer fmt.Println("Deferred x:", x) // x is 10 here
  x = 20
  fmt.Println("Current x:", x)
}
