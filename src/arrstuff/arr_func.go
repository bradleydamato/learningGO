package main 
//testing a thesis on as to how arrays and slices work in Go 
import "fmt"


var q []int = []int{1,2,3}

var r []int = []int{4,5}

func main(){
	n := []int{}
	for i := 0; i <= len(q)-1; i++ {
		n = append(n,q[i])
			}
	fmt.Println(n)
}
