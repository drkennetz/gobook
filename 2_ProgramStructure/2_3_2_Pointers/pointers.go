package main

import "fmt"

func main() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // 1
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // 2
	// each component of an aggregate type - a field of a struct or an element of an array - is also a variable
	// thus having an address too.

	// the zero-value of a pointer of any type is nil. The test:
	// p != nil is true if p points to a variable.
	// Pointers are comparable. They are equal if and only if they point to the same variable or both are nil

	var j, k int
	fmt.Println(&j == &j, &j == &k, &j == nil) // true, false, false
	var cool = f()
	// each call of f returns a distinct value
	fmt.Println(f() == f()) // false
	fmt.Println(cool)

	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3"
}

// variable v inside f will remain in existence even after the call has returned, and the pointer cool will still refer to it
func f() *int {
	v := 1
	return &v
}

// because a pointer contains the address of a var, passing a pointer argument to a function makes it possible for the function to update the variable that was indirectly passed.
// For example, this function increments the variable that its arguments points to an returns the new value of the variable so it may be used in an expression

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}
