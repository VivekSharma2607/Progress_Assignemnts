package main 
import "fmt"
func push (queue[] int , temp int)[] int{
	fmt.Print("The element that is being added to the queue = " , temp)
	fmt.Print("\n")
	queue = append(queue, temp)
	return queue
}
func pop(queue[] int ) [] int {
	temp := queue[0]
	fmt.Print("The deleted element from the queue = " , temp)
	fmt.Print("\n")
	return queue[1:]
}
func main() {
	var queue[] int ;
	queue = push(queue , 10)
	queue = push(queue , 20)
	queue = push(queue , 30)
	fmt.Print(queue)
	fmt.Print("\n")
	queue = pop(queue)
	fmt.Print(queue)
	fmt.Print("\n")
}