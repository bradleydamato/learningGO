package main 
//testing a thesis on as to how arrays and slices work in Go 
import "fmt"


var q []int = []int{1,2,3}

var r []int = []int{4,5}



func main(){
	//SHORTEST WAY
	w := append(q[:],r[:]...)

	//SHORTER WAY
	x :=  append(q[:])
	x =  append(x,r[:]...)

	//LONG WAY
	n := []int{}
	for i := 0; i <= len(q)-1; i++ {
		n = append(n,q[i])
			}
	for i := 0; i <= len(r)-1; i++ {
		n = append(n,r[i])
			}
	fmt.Println(n)
	fmt.Println(x)
	fmt.Println(w)
	test := reverse(w)
	fmt.Println(test)
}

//reverse slice of ints in place 
func reverse(s []int) []int {
	for i, j := 0,len(s)-1; i < j; i, j = i+1, j-1{
		s[i],s[j] = s[j],s[i]
	}
	return s
}
