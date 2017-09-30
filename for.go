package main

import "fmt"

func main() {
  i := 1

  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("loop")
    break
  }

  for n := 0; n <= 5; n++ {
    if n % 2 == 0 {
      continue
    }
    fmt.Println(n)
  }

  // Ping-Pong
  for ping_pong := 0; ping_pong <= 30; ping_pong++ {
    if ping_pong % 15 == 0 {
      fmt.Println("ping-pong")
      continue
    }
    fmt.Println(ping_pong)
  }
}
