package main

import "fmt"

func swap(x *int, y *int) () {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	x := 1
	y := 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
