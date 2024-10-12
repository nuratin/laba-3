package main

import "fmt"

func addStars(s string) string {
	result := ""
	for i, c := range s {
		result += string(c)
		if i < len(s)-1 {
			result += "*"
		}
	}
	return result
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(addStars(input))
}
