package main

import "fmt"

func main() {
	defer fmt.Println("world")   //6th to execute
	defer fmt.Println(double(1)) //5th to execute
	defer fmt.Println(double(2)) //4th
	defer fmt.Println(double(3)) //3rd
	defer fmt.Println(double(4)) //2nd
	defer fmt.Println(double(5)) //1st


	fmt.Println("hello")
}

func double(x int) int {
	return x + x
	}
