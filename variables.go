package main

import "fmt"

func main() {
    var a string = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true // Go understands true is an implied type boolean
    fmt.Println(d)

    var e int // An integer is 0 when initialized but not assigned a value
    fmt.Println(e)

    f := "short" // Shorthand for initializing and declaring variable
    fmt.Println(f)
}
