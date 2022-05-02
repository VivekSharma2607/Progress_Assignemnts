package main

import (
	"fmt"
	// "internal/fmtsort"
)

func main(){
	var a  , rev int
	fmt.Println("Enter the number you want to check = ")
	fmt.Scanf("%d" , &a)
	rem := 0
	for i := a ; i > 0 ; i = i/10 {
		rem = i % 10
		rev = rev*10  + rem
	}
	if rev == a{
		fmt.Println("The number is palindrome")
	}else{
		fmt.Println("The number is not a palindrome")
	}
}