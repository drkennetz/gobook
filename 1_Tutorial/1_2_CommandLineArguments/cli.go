package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// explicit declaration to 0 value ("") for string, in a loop
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Printing from for loop: ", s)

	// short declaration in a range
	s2, sep2 := "", ""
	for _, arg := range os.Args[0:] {
		s2 += sep2 + arg
		sep = " "
	}
	fmt.Println("Printing from range: ", s2)

	// string concat is expensive, so if the data is large, we should join
	fmt.Println("Printing from strings.Join(): ", strings.Join(os.Args[0:], " "))

	//finally, if we don't care about the format and want to print debug:
	fmt.Println("Printing to debug: ", os.Args[0:])
}

// exercise: benchmark for loop, range loop, and join
func ForLoop(args []string) string {
	s := ""
	sep := ""
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func RangeLoop(args []string) string {
	s := ""
	sep := ""
	for _, v := range args {
		s += sep + v
		sep = " "
	}
	return s
}

func Join(args []string) string {
	return strings.Join(os.Args, " ")
}