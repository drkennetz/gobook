package main

/*
The lifetime of a var is the inerval of time during which it exists as the program executes
the lifetime of a package level var is the entire execution of the program
local variables have dynamic lifetimes. They stay around until they become unreachable at which point its storage may be recycled
function params and results are local vars too; they are created each time their enclosing func is called
if no path exists to a var (either by pointers or other kinds of references), the var is unreachable.
a local variable may outlive a single iteration of an enclosing loop, if it is allocated in such a way that it is still reachable
*/

// consider this example
var global *int

func f() {
	var x int
	x = 1
	global = &x
}

// x must be heap allocated because it is still reachable from the variable global after f has returned

func g() {
	y := new(int)
	*y = 1
}

// conversely, when g returns, the variable *y becomes unreachable and can be recycled.
// Since *y does not escape from g, it's safe for the compiler to allocate *y on the stack
// even though it was allocated with new

// Garbage collection is a tremendous help in writing correct programs, but it does not relieve you of the burden
// of thinking about memory. You don't need to explicitly allocate and free memory, but to write efficient programs
// you still need to be aware of the lifetime of vars
// for example, keeping unnecessary pointers to short-lived objects within long-lived objects, especially global vars
// will prevent the garbage collector from reclaiming short-lived objects
