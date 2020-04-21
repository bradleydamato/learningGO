package main

import "fmt"

type node struct {
	value       int
	left, right *node
}

//Sort sorts values
func Sort(values []int) {
	var root *node
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)

}

//appendValues appends the elements of t to values in order
// and returns the resultign slice
func appendValues(values []int, t *node) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *node, value int) *node {

	if t == nil {
		//equivalent to return &node{value: value}
		t = new(node)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)

	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	test := []int{5, 23, 6, 2342, 62, 432, 3, 634, 7453, 4}
	fmt.Println(test)
	Sort(test)
	fmt.Println(test)
}
