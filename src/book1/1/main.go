package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args) < 2 {
		fmt.Println("Please set a file name.")
		return
	}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {

		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s %d\n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	imput := bufio.NewScanner(f)
	fmt.Println()
	for imput.Scan() {
		counts[imput.Text()]++
	}
}
