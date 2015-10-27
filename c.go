package main

import (
    "fmt"
    "time"
)

func thread(name string, messages chan string) {
    for i:=0; i < 10; i++ {
        fmt.Println("[", name, "] ", i)
        time.Sleep(1000 * time.Millisecond)
    }
    messages <- "done"
}


func main() {
    messages := make(chan string)
    go thread("a", messages)
    go thread("b", messages)
    go thread("c", messages)
    go thread("d", messages)
    var input string

    for i:=0; i < 4; i++ {
        input = <- messages
        fmt.Println(input)
    }
}
