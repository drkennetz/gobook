package tuple_assignment

// tuple assignment allows several vars to be assigned at once
// all of the right handed expressions are evalueated before any of the variables are updated
// making this form most useful when some of the vars appear on both sides of the assignment
// examples:

func GreatestCommonDivisor(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func NFibonnaci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

/*
map lookup, type assertion, or channel receive appears in an assignment in which two results are expected,
each produces a boolean result:

    v, ok = m[key] // map lookup
    v, ok = x.(T) // type assertion
    v, ok = <-ch // channel receive

As with variable declarations, we can assign unwanted values to the blank identifier:

    _, err = io.Copy(dst, src) // discard the byte count
    -, ok = x.(T) // check the type but discard the result
*/
