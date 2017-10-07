package main

import "fmt"

func main() {
  // Fun story. Maps are Go's version of Hashes/Dictionaries
  m := make(map[string]int) // Inits a map of String:Int
  // make(map[key-type]val-type) .... All these type declarations, I've been using too many scripting languages for this haha

  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m)

  v1 := m["k1"]
  fmt.Println("v1: ", v1)

  fmt.Println("len:", len(m))

  delete(m, "k2")
  fmt.Println("map:", m)

  _, prs := m["k2"] // The blank identifier allows us to check the presence of a key (prs) without getting it mixed up with a keys value
  fmt.Println("prs:", prs)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
}
