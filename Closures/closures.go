package main

import "fmt"

func main() {
    // Outer function
    counter := 0
    increment := func() int { // Closure
        counter++
        return counter
    }

    fmt.Println(increment()) // Output: 1
    fmt.Println(increment()) // Output: 2
    fmt.Println(increment()) // Output: 3
}
