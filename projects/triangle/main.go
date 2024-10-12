package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	hypotenuse := math.Sqrt(a*a + b*b)
	fmt.Println(hypotenuse)
}
