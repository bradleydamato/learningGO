package main 

import "fmt"

func remove(slice []int, i int) []int{
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice)-1]
	}
func main() {
	s := []int{0,1,2,3}
	fmt.Println(remove(s,2))//removes at index 2, not the element named 2 
}
