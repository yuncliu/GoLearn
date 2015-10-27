package main

import (
	"fmt"
)

func main() {
    var s []int //s is a slice
    s = append(s, 1)
    s = append(s, 2)
    s = append(s, 3)
    print_slice(s)
    f(s)
    print_slice(s)
}

func f(s []int) {  // value can be changed, but lenght can't
    for i:= 0; i < len(s); i++ {
        s[i] *=10
    }
    s = append(s, 100)  //this element can't get out
    print_slice(s)
}

func print_slice(s []int) {
    fmt.Println("------------")
    for i:= 0; i < len(s); i++ {
        fmt.Println(s[i])
    }
}
