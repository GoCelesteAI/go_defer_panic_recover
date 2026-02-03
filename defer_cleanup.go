package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println("=== Defer for Cleanup ===")

  // Common pattern: open resource, defer close
  writeFile()

  fmt.Println("\n=== Defer with Named Returns ===")
  result := namedReturn()
  fmt.Println("Result:", result)

  fmt.Println("\n=== Defer in Loops ===")
  deferInLoop()
}

func writeFile() {
  file, err := os.Create("test.txt")
  if err != nil {
    fmt.Println("Error creating file:", err)
    return
  }
  defer file.Close() // Always closes, even if error below

  _, err = file.WriteString("Hello, Go!\n")
  if err != nil {
    fmt.Println("Error writing:", err)
    return
  }
  fmt.Println("File written successfully")

  // Cleanup the test file
  defer os.Remove("test.txt")
}

func namedReturn() (result int) {
  defer func() {
    result = result * 2 // Modifies the return value
  }()
  return 5 // Returns 10, not 5
}

func deferInLoop() {
  // Be careful with defer in loops
  for i := 0; i < 3; i++ {
    defer fmt.Println("Loop defer:", i)
  }
  fmt.Println("Loop done")
  // All defers run when function exits, not after each iteration
}
