package main

import "fmt"


func main() {
	result := add(3,5)
	fmt.Println("The result is:", result)
}

func add(a int, b int) int {
	return a + b
}