package main

import (
	"fmt"
	"time"
)

func thread(name string, messages chan string) {
	for {
		x := <-messages
		fmt.Println(name, "got ", x)
	}
}

func main() {
	messages := make(chan string)
	go thread("1", messages)
	go thread("2", messages)
	go thread("3", messages)
	for i := 0; i < 10; i++ {
		messages <- "hellow"
		time.Sleep(1000 * time.Millisecond)
	}
}
