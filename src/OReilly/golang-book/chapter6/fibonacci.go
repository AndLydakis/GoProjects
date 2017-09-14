package main

import "fmt"

func fib(x int) int {
	if x == 1 || x == 0 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(fib(i))
	}
}
