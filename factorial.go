package main

import (
    "fmt"
)

func main() {
    fmt.Print("factorial = ", factorial(10))
}

func factorial(a int) int {
    return fact(a, 1)
}

/*tail recursion*/
func fact(a int, b int) int {
    if a == 1 {
        return b
    }
    return  fact(a-1, a*b)
}
