package main

import (
	"fmt"
)

func main() {
    m := make(map[int]int)
    m[1] = 10
    m[2] = 20
    fmt.Println(m)
    delete(m, 2)
    fmt.Println(m)
}
