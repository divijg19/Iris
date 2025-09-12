// This Go file offers an in-depth exploration of advanced and nuanced topics in the Go programming language.
// It is designed for individuals who have a foundational understanding of Go and are looking to deepen their expertise.
// We will delve into concurrency, interfaces, and error handling, providing detailed explanations and practical examples.

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// SECTION 1: ADVANCED CONCURRENCY IN GO

// Go is renowned for its powerful and simple approach to concurrent programming. [1]
// Concurrency is the ability to have multiple tasks in progress at the same time. [10]
// This is achieved through goroutines and channels, which are core features of the language. [9]

// Goroutines are lightweight threads managed by the Go runtime. [1]
// They are significantly cheaper than traditional threads, and it's common to have thousands or even millions of them running simultaneously. [10]

// Channels provide a way for goroutines to communicate and synchronize their execution. [2]

// A simple function to demonstrate a goroutine.
func printMessage(message string, wg *sync.WaitGroup) {
	// Decrement the WaitGroup counter when the goroutine completes.
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

// Another function to illustrate concurrent execution.
func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

// A function that simulates checking the status of a website.
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}

// SECTION 2: IN-DEPTH LOOK AT INTERFACES

// Interfaces in Go provide a way to specify the behavior of an object. [9]
// An interface is a type that defines a set of method signatures. [2]
// A type implements an interface by implementing its methods. There is no explicit "implements" keyword. [6]

// Define an interface for geometric shapes.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct for a rectangle.
type Rectangle struct {
	Width, Height float64
}

// Implement the Shape interface for Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Define a struct for a circle.
type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// A function that can take any type that implements the Shape interface.
func PrintShapeDetails(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

// The empty interface `interface{}` is a special case.
// It has no methods, so all types implement it. This allows for creating functions that can accept any type.
func PrintAnything(v interface{}) {
	fmt.Println("The value is:", v)

	// We can use a type switch to inspect the underlying type of an empty interface.
	switch v.(type) {
	case int:
		fmt.Println("This is an integer.")
	case string:
		fmt.Println("This is a string.")
	default:
		fmt.Println("This is some other type.")
	}
}

// SECTION 3: ROBUST ERROR HANDLING

// Go's approach to error handling is explicit and straightforward. [9]
// Functions that can fail return an error as the last return value.
// The caller is expected to check the error value.

// A custom error type.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// A function that might return an error.
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	// --- ADVANCED CONCURRENCY IN ACTION ---

	// A WaitGroup waits for a collection of goroutines to finish.
	var wg sync.WaitGroup

	// Add two to the WaitGroup counter.
	wg.Add(2)

	// Start two goroutines.
	go printMessage("Hello from a goroutine!", &wg)
	go printMessage("This is another goroutine.", &wg)

	// Wait for the goroutines to finish.
	wg.Wait()

	fmt.Println("All goroutines finished.")

	// Example of channels for communication.
	go count("sheep")
	go count("fish")

	// The `time.Sleep` is a simple way to wait for the goroutines to finish.
	// In a real application, you would use more robust synchronization mechanisms like WaitGroups or channels.
	time.Sleep(time.Second * 3)

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel to communicate the status of the websites.
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Wait for the goroutines to finish and print the results.
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	// --- IN-DEPTH INTERFACES IN ACTION ---

	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	fmt.Println("Rectangle details:")
	PrintShapeDetails(rect)

	fmt.Println("\nCircle details:")
	PrintShapeDetails(circ)

	fmt.Println("\nDemonstrating the empty interface:")
	PrintAnything(42)
	PrintAnything("hello")
	PrintAnything(true)
	PrintAnything(rect)

	// --- ROBUST ERROR HANDLING IN ACTION ---

	// Simple error handling.
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// Using the custom error type.
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("End of the program.")
}
