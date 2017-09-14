package main

import "fmt"

func main() {

	add := func(args ... int) int {
		sum := 0
		for _, v := range args {
			sum += v
		}
		return sum
	}

	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(add(1, 2, 3, 4, 5, 6))
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(add(x...))
}
