//Go does not have a set type, which is interesting and different from Python3.
//One does wonder why anyone uses go
//Not to say I don't Enjoy it, but, I mean I'll be frank in saying I have no clue if it is faster than Python3.
//I'm going to google that right now.

//OKAY GO is substantially faster than Python3 on several metrics. Go is slower than C++ which I expected. Go is == Java in  some respects but slower in others.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) //set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedupe %v\n", err)
		os.Exit(1)
	}

}
