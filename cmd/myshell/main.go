package main

import (
	"bufio"
	"fmt"

	// Uncomment this block to pass the first stage
	// "fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
}
