package main 

import "fmt"

func main() {
	for i := 3; i > -3; i--{
		
		x := 3/i
		fmt.Println(x)


			defer func() {
        			if p := recover(); p != nil {	
				fmt.Println("STOP TRING TO DIVIDE BY ZERO!!!!")}}()
		}}
