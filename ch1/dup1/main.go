// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

//Added by Eric:
//[eric@linux-x2vq dup1](master)$ echo 'this is hello'   > /tmp/lines
//[eric@linux-x2vq dup1](master)$ echo 'this is filter1' >> /tmp/lines
//[eric@linux-x2vq dup1](master)$ echo 'this is filter2' >> /tmp/lines
//[eric@linux-x2vq dup1](master)$ echo 'this is filter1' >> /tmp/lines
//[eric@linux-x2vq dup1](master)$ echo 'this is hello'   >> /tmp/lines
//[eric@linux-x2vq dup1](master)$ echo 'this is hello'   >> /tmp/lines
//[eric@linux-x2vq dup1](master)$ cat /tmp/lines | go run main.go

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
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
