package main

import "fmt"
import "math"

func Square(n int) string {
	s := ""
	for n > 0 {
		s = fmt.Sprint(math.Pow(float64(n%10), 2)) + s
		n /= 10
	}

	return s
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Print(Square(num))
}
