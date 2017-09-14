package main

import "fmt"

func great(args ...int) int {
	g := 0
	for _, v := range args {
		if v > g {
			g = v
		}
	}
	return g
}
func main() {
	fmt.Println(great(1, 2, 3, 4, 5, 6, 7, 8))
}
