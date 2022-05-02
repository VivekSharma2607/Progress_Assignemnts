package main

import "fmt"

func help(n int) int {
	mp := []int {0,1}

	for i:= 3 ; i <= n ; i++ {
		mp1 := mp[0] + mp[1]
		mp = []int {mp[1] , mp1}
	}
	if n > 1 {
		return mp[1]
	}
	return mp[0]
}

func main() {
	var n int
	fmt.Printf("Enter the element number you need = ")
	fmt.Scanf("%d" , &n)
	var ans int
	ans = help(n)
	fmt.Print("The nth element of the series = " , ans)

}