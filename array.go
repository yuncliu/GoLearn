package main

import (
	"fmt"
)

type xxx []float64

func main() {
	var a [5]float64
	for i := 0; i < len(a); i++ {
		a[i] = float64(i)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	var c xxx
	c = a[:] // c is a slice
	fmt.Println("Average = ", average(c))
	copy(a[:], c) // copy slice to array,use [:] to trick copy a is a slice
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println(c.Aver())
}

func average(array []float64) float64 {
	sum := 0.0
	for i := 0; i < len(array); i++ {
		sum += array[i]
		array[i] *= 10
	}
	return sum / float64(len(array))
}

func (array *xxx) Aver() float64 {
	fmt.Println("xxxx")
	for i := 0; i < len(*array); i++ {
		fmt.Println("xxxxxxxxx = ", (*array)[i])
	}
	return 100.0
}
