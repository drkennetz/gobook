package duplines

import (
	"bufio"
	"fmt"
	"os"
)

/*
Modify dup2 to print the names of all files in which each duplicated line occurs
*/

func Dup2PrintName() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines2(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines2(f, counts)
			f.Close()
		}
	}
	for fname := range counts {
		for _, n := range fname {
			if n > 1 {
				fmt.Printf("%s\n", fname)
			}
		}
	}
}

func countLines2(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}
