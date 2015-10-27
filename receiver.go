package main

import (
	"fmt"
)

/*Cannot define new methods on non-local type, so define a local type */
type XXX int;

func main() {
    var a XXX
    a = 10
    fmt.Println(a.time10())
    a.time100()
    fmt.Println(a.time10())
}

func (a *XXX) time10() int {
    return int((*a) * 10)
}

func (a *XXX) time100(){
    *a *=100
}
