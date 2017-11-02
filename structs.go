package main

import "fmt"

type person struct {
  name string
  age int
}

// My very own cryptocurrency struct!!
type gocoin struct {
  storage_address string
  value float32
  history []string
}

func main() {
  fmt.Println(person{"Bob", 20})

  fmt.Println(person{name: "Bob", age: 30})

  fmt.Println(person{name: "Fred"})

  fmt.Println(&person{name: "Ann", age: 40})

  s := person{name: "Dawson", age: 22}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 55
  fmt.Println(sp.age)

  fmt.Println(gocoin{storage_address: "sdkfh4h3l2h5l", value: 4.7})
  fmt.Println(gocoin{"7dsf78a90sdfh", 11.3, []string{"sdkfh4h3l2h5l", "45jh43kjlh4hh"}})
}
