package main

import "fmt"

type Point struct {
	x int 
	y int
}

func main() {
	var p1 Point = Point{1,2}
	p1.x = 10
	fmt.Println(p1.x)
}