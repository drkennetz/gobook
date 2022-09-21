package main

import "fmt"

// new returns a pointer to the object it creates.
// It is syntactic sugar that allows you to init a val as a pointer without a dummy variable name
// each call to new returns a distinct var with a unique address

func main() {
	fmt.Println(newIntNew())
	fmt.Println(newIntVar())
}

// Utilizes the new function to return a pointer to an int
func newIntNew() *int {
	return new(int)
}

func newIntVar() *int {
	var dummy int
	return &dummy
}
