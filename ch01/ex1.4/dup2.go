// Dup2 prints the filenames that has duplicate lines
// It reads from stdin or from a list of named files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var res []string
	if len(files) == 0 {
		res = countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			res = countLines(f, counts, arg)
			f.Close()
		}
		fmt.Println(res)
	}
}

func countLines(f *os.File, count map[string]int, arg string) []string {
	input := bufio.NewScanner(f)
	var res []string
	for input.Scan() {
		count[input.Text()]++
		if count[input.Text()] > 1 {
			res = append(res, arg)
		}
	}
	return res
}
