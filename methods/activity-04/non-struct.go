package main

import (
	"fmt"
	"math"
)

type myFloat float64

func (m myFloat) ceil() float64 {
	return math.Ceil(float64(m))
}

func main() {
	num := myFloat(0.9)
	fmt.Println(num.ceil())
}
