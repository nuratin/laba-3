package main

import "fmt"

func maxDigit(s string) byte {
	max := byte('0')
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(string(maxDigit(input)))
}
