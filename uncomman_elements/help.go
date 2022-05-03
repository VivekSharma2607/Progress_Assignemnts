package main

import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}
	b := [5]int{2,3,0,1,5}
	mp := make(map[int]int)
	for _,i := range a {
		if mp[i] == 0{
			mp[i] = 1
		}else {
			mp[i] += 1
		}
	}
	for _,i := range b {
		if mp[i] == 0{
			mp[i] = 1
		}else {
			mp[i] += 1
		}
	}
	for i , value := range mp {
		if value == 1{
			fmt.Printf("The eleents occuring once in both the arrays  = %d" , i)
			fmt.Print("\n")
		}	
	}
}