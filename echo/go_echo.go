package main 

import (
"fmt"
"os"
"time"
)


func main() {
	start := time.Now()
	var s, sep string 
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
		fmt.Println(i,os.Args[i])
	}
	fmt.Println(s)
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println(diff)



}
