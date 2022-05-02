package main

import "fmt"

func help (a string , b string) bool {
	len1 := len(a)
	len2 := len(b)
	if len1 != len2 {
		return false
	}
	mp1 := map[byte]int{}
	mp2 := map[byte]int{}

	for i:= 0 ; i < len1 ; i++{
		mp1[a[i]] +=1
		mp2[b[i]] += 1
	}
	for i:= 0 ; i < len1 ; i++{
		if mp1[a[i]] != mp2[b[i]]{
			return false
		}
	}
	return true
}
func main(){
	var a , b string

	fmt.Println("Enter the first string")
	fmt.Scan(&a)
	fmt.Println("Enter the second string ")
	fmt.Scan(&b)
	var ans bool = false
	ans = help(a , b)
	if ans == false {
		fmt.Println("The strings are not anagram")
	}else {
		fmt.Println("The strings are anagram")
	}

}