// Every Go program starts with a package declaration.
// The 'main' package is special; it tells the Go compiler
// that the program is executable. [3, 4]
package main

// The 'import' keyword is used to include code from other packages.
// The 'fmt' package (short for format) provides functions for
// formatted I/O, like printing to the console. [3, 4]
import "fmt"

// The 'main' function is the entry point of the program.
// Execution starts here. [4]
func main() {
    // 'Println' is a function from the 'fmt' package that prints
    // a line of text to the console. [4]
    fmt.Println("Hello, World!")
}
