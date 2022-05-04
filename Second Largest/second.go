package main

import "fmt"

 

func main() {
	var n int
	fmt.Printf("Enter the size of array = ")
	fmt.Scanf("%d" , &n)
	fmt.Print("Start entering the value in an array \n")
	var a = make([] int, n)
	for i := 0 ; i < n ; i++ {
		fmt.Scanf("%d" , &a[i])
	}
	
	largest := 0 
	second := 0
	for i:= 0 ; i < n ; i++{
		if a[i] > largest {
			second = largest
			largest = a[i]
		}else if a[i] > second && a[i] != largest {
			second = a[i]
		}
	}


	fmt.Printf("The largest element of an array = %d " , largest)
	fmt.Print("\n")
	fmt.Printf("The second largest element of an array = %d" , second)
	fmt.Print("\n")

}
