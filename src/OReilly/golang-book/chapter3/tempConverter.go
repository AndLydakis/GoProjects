package main

import "fmt"

func main() {
	var farenheit float64
	fmt.Println("Enter farenheit temperature ")
	fmt.Scanf("%f", &farenheit)
	fmt.Println((farenheit - 32) * (5 / 9))
}
