package main

import "fmt"

var slicer =  []int{} 

func main() {
fmt.Println(slicer) 
slicer = append(slicer,1)
fmt.Println(slicer) 

slicer = append(slicer,2)
slicer = append(slicer,3)
slicer = pop(slicer)
fmt.Println(slicer) 



}
func pop(stack []int) []int {
	fmt.Println(stack[len(stack)-1]) //prints top element
	stack = stack[:len(stack)-1] //gets rid of top element
	return stack

}
