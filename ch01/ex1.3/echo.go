package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Time 1
	start := time.Now()
	s, sep := "", ""
	for i := 1; i < 100000; i++ {
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
	fmt.Println(time.Since(start).Seconds())

	// Time 2
	start = time.Now()
	for i := 1; i < 100000; i++ {
		strings.Join(os.Args[1:], " ")
	}
	fmt.Println(time.Since(start).Seconds())
}
