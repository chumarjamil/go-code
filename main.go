package main

import "fmt"


func main() {
	// loop in golang

	for i :=0; i < 5; i++ {
		fmt.Println(i)
	}

	// Looping through it
	nums := [] int{1,2,3,4,5}
	for index, value := range nums {
		fmt.Println("Index", index, "Value", value)
	}
}