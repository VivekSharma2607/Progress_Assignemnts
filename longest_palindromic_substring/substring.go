package main 
import "fmt"

func main() {
	
	var a string 
	fmt.Print("Enter the string = ")
	fmt.Scanf("%s" , &a)
	start := -1
	maxlenght := 0

	for i:= 0 ; i < len(a) ; i++{
		low := i-1 
		high := i+1
		for low >= 0 && a[i] == a[low] {
			low--
		}
		for high < len(a) && a[i] == a[high] {
			high++
		}
		for low >= 0 && high < len(a) && a[low] == a[high] {
			low--
			high++
		}
		len := high-low-1
		if maxlenght < len {
			start = low+1
			maxlenght = len
		}
	}
	var ans string
	ans = a[start : start+maxlenght]
	fmt.Println("The longest Palindromic substring = ",ans)
	fmt.Println("The lenght of the longest Palindromic substring = ",maxlenght)
}