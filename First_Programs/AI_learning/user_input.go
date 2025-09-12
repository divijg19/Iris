package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a new reader that reads from standard input (the keyboard). [14]
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")

	// ReadString reads input until the first occurrence of a delimiter.
	// In this case, it's the newline character '\n', which is added
	// when you press Enter. [14]
	name, _ := reader.ReadString('\n')

	// The input includes the newline character, so we remove it.
	name = strings.TrimSpace(name)

	// Greet the user with their name.
	fmt.Printf("Hello, %s!\n", name)
}
