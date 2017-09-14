package main

import "fmt"

func half(x int) (int, bool) {
	if (x % 2) == 0 {
		return x / 2, true
	} else {
		return x / 2, false
	}
}

func main() {
	fmt.Println(half(5))
	fmt.Println(half(4))
}
