package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Printf("ala ma kotka2")
	// NOTE: Ignoring potential errors form input.Error()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d, %s\n", n, line)
		}
	}
}
