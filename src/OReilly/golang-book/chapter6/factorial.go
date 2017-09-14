package main

import "fmt"

func factorial(i uint) uint {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	fmt.Println(factorial(0))
	fmt.Println(factorial(1))
	fmt.Println(factorial(5))
}
