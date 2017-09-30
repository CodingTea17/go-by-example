package main

import "fmt"
import "math"

const s string = "constant"

func main() {
  fmt.Println(s)

  const n = 500000000

  const d = 3e20 / n // A numeric constant w/o a type
  fmt.Println(d)

  fmt.Println(int64(d)) // Explicitly given a type of int64

  fmt.Println(math.Sin(n)) //Contextually given a float64 type
}
