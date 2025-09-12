package main

import "fmt"

func main() {
    // Declare a string variable and initialize it.
    var greeting string = "Hello from a variable!"
    fmt.Println(greeting)

    // Go can often infer the type of a variable.
    // The ' := ' syntax is a shorthand for declaring and initializing.
    name := "Go"
    year := 2025

    // Print formatted strings using 'Printf'.
    // '%s' is a placeholder for a string.
    // '%d' is a placeholder for an integer.
    // '\n' creates a new line.
    fmt.Printf("Welcome to %s in the year %d!\n", name, year)
}
