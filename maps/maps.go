// Implementing maps in golang 


package main 

import "fmt"

func main () {
	var mp map[int]int = map[int]int {
		1 : 5,
		2 : 20 ,
	}
	for key , value := range mp {
		fmt.Printf("%d : %d" , key , value)
		fmt.Print("\n")
	}
}