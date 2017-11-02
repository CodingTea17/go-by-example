package main

import "fmt"

func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i)

  // i is 1 and when we pass it as an argument
  zeroval(i)
  // it's value isn't changes
  fmt.Println("zeroval:", i)

  // i is 1 and when we pass it by reference (it's actual spot in memory)
  zeroptr(&i)
  // it's value becomes 0 as dictated by the zeroptr func
  fmt.Println("zeroptr:", i)

  // The memory address of i 
  fmt.Println("pointer:", &i)
}
