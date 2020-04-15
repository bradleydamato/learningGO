package main 

import (
"fmt"
"os"
"strings"
"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], ""))
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println(diff)
}