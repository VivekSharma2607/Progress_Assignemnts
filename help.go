package main

import "fmt"

func check (i interface{} , j interface{}) {
	if(i == j){
		fmt.Println("Is equal")
	}else {
		fmt.Println("Not Equak")
	}
}


func main() {
	var i interface{} = 90
	var j interface{}= 89

	check(i,j)

	var a interface {} = "Hlo"
	var b interface {} = "Hlo"
	check(a,b)
}