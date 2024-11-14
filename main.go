package main

import (
	"fmt"
	"time"
)

// This code starts sayHello in a separate goroutine, so the main program doesn’t wait for it to finish.
func sayHello() {
	fmt.Println("Hello, Goroutine")
}

func main() {
	// loop in golang

	// for i :=0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// // Looping through it
	// nums := [] int{1,2,3,4,5}
	// for index, value := range nums {
	// 	fmt.Println("Index", index, "Value", value)
	// }

	// Pointers in golang
	// x := 10
	// p := &x
	// fmt.Println(*p)

	go sayHello() // start a new go routine
	time.Sleep(time.Second) // wait for the seconds the go routine compelte
}

