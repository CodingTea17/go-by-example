package main

import "fmt"

// Simple recursion that
func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}

func main() {
  fmt.Println(fact(4));
  // 4 * (3 * (2 * (1 * (1))))
}
