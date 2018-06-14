package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//$ echo 'hello'   > /tmp/lines
//$ echo 'filter' >> /tmp/lines
//$ cat /tmp/lines | go run line-filters.go

func main() {

	// Wrapping the unbuffered `os.Stdin` with a buffered
	// scanner gives us a convenient `Scan` method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` returns the current token, here the next line,
		// from the input.
		ucl := strings.ToUpper(scanner.Text())

		// Write out the uppercased line.
		fmt.Println(ucl)
	}

	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
