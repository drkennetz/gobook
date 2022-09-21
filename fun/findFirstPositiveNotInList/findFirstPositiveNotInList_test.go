package main

import (
	"fmt"
	"testing"
)

func TestFindFirstPositiveNotInList(t *testing.T) {
	type Case struct {
		input  []int
		result int
	}
	cases := []Case{
		{input: []int{3, 4, -1, 1}, result: 2},
		{input: []int{1, 2, 0}, result: 3},
		{input: []int{0}, result: 1},
		{input: []int{}, result: 1},
	}
	for i, v := range cases {
		if v.result != FindFirstPositiveNotInList(v.input) {
			fmt.Println("Failed test:", i)
		}
	}
}
