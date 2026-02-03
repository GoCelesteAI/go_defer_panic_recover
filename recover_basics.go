package main

import "fmt"

func main() {
  fmt.Println("=== Recover Basics ===")

  // Recover allows graceful handling of panics
  fmt.Println("Before risky operation")
  riskyOperation()
  fmt.Println("After risky operation - program continues!")

  fmt.Println("\n=== Practical Example ===")
  results := processItems([]int{10, 5, 0, 8, 2})
  fmt.Println("Results:", results)
}

func riskyOperation() {
  // Defer with recover must be in the same function
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered from:", r)
    }
  }()

  fmt.Println("About to panic...")
  panic("something went wrong!")
  fmt.Println("This never runs")
}

func processItems(items []int) []string {
  var results []string

  for _, item := range items {
    result := safeProcess(item)
    results = append(results, result)
  }

  return results
}

func safeProcess(n int) string {
  defer func() {
    if r := recover(); r != nil {
      // Silently recover, return will use default
    }
  }()

  if n == 0 {
    panic("cannot process zero")
  }

  return fmt.Sprintf("100/%d=%d", n, 100/n)
}
