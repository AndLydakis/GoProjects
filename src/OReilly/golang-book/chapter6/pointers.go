package main

import "fmt"

func zeroX(x uint) {
	fmt.Println(x)
	x = 0
}

func zeroXPtr(x *uint) {
	fmt.Println(x)
	*x = 0
}

func main(){
	 var x uint
	 x = 5
	 zeroX(x)
	 fmt.Println(x)

	 xPtr := new(uint) //returns pointer
	 *xPtr = 5
	 fmt.Println(*xPtr)
	 zeroXPtr(xPtr)
	 fmt.Println(*xPtr)
}