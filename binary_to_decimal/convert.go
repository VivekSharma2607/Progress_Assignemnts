package main

import (
	"fmt"
	"math"
)

func convert(n int) int {
	var rem int
	i := 0
	decimalnum := 0
	for n != 0 {
		rem = n%10
		n /= 10
		decimalnum += rem * int(math.Pow(2,float64(i)))
		i++
	}
	return decimalnum
}

func main() {
	var a int 
	fmt.Printf("Enter the binary number  to be converted = ")
	fmt.Scanf("%d" , &a)
	fmt.Println(convert(a))
}