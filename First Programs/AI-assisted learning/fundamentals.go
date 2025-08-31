// This file is designed to be a comprehensive, in-depth tutorial of the Go programming language.
// It is intended for beginners with some programming experience. [2]
// We will cover the fundamental concepts of Go, starting from the very basics and moving to more advanced topics.
// Each concept is explained with comments and illustrated with code examples.

// Every Go program starts with a package declaration. [9]
// The 'main' package is the entry point for executable programs. [10]
package main

// The 'import' statement is used to include packages that contain functions and features you want to use. [9]
// 'fmt' is a standard Go package for formatted I/O (input/output). [10]
import (
	"errors"
	"fmt"
	"math"
	"time"
)

// The 'main' function is the entry point of the program. When you run your program, this is the first function that gets executed. [9]
func main() {
	// Let's start with the traditional "Hello, World!".
	// fmt.Println prints a line of text to the console. [10]
	fmt.Println("Hello, World!")

	// 1. VARIABLES AND DATA TYPES
	// Variables are used to store and manage data in a program. [15]
	// In Go, you can declare a variable using the 'var' keyword, followed by the variable name and its type. [8]
	var anInteger int = 42
	fmt.Println("This is an integer:", anInteger)

	// Go also supports a short variable declaration using ':=', which infers the type of the variable. [8]
	aString := "This is a string."
	fmt.Println(aString)

	// Go has several basic data types:
	// - bool: represents a boolean value, either true or false. [15]
	var aBoolean bool = true
	fmt.Println("This is a boolean:", aBoolean)

	// - int, int8, int16, int32, int64: signed integers of different sizes.
	// - uint, uint8, uint16, uint32, uint64, uintptr: unsigned integers.
	// - float32, float64: floating-point numbers. [15]
	aFloat := 3.14
	fmt.Println("This is a float:", aFloat)

	// - string: a sequence of characters. [15]
	// - rune: a single character (an alias for int32).
	// - byte: an alias for uint8.

	// You can also declare multiple variables at once.
	var (
		x int    = 10
		y string = "hello"
	)
	fmt.Println("Multiple variables:", x, y)

	// Constants are like variables, but their values cannot be changed.
	const PI = 3.14159
	fmt.Println("The value of PI is:", PI)

	// 2. CONTROL FLOW
	// Control flow statements direct the execution of a program. [8]

	// 'if' and 'else' statements are used for conditional logic.
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// You can also have 'else if'.
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// The 'for' loop is the only looping construct in Go.
	// It has three components separated by semicolons:
	// - the init statement: executed before the first iteration
	// - the condition expression: evaluated before every iteration
	// - the post statement: executed at the end of every iteration
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// The 'init' and 'post' statements are optional. This creates a while-like loop.
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("Sum is:", sum)

	// 'switch' is a convenient way to write multiple 'if-else' statements.
	day := "Wednesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("Invalid day.")
	}

	// 3. ARRAYS, SLICES, AND MAPS

	// An array has a fixed size.
	var anArray [5]int
	anArray[0] = 10
	fmt.Println("First element of the array:", anArray[0])

	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// Slices are much more common in Go than arrays.
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("This is a slice:", aSlice)
	aSlice = append(aSlice, 6)
	fmt.Println("Appended slice:", aSlice)

	// A map stores key-value pairs.
	aMap := make(map[string]int)
	aMap["apple"] = 1
	aMap["banana"] = 2
	fmt.Println("This is a map:", aMap)
	fmt.Println("The value for 'apple' is:", aMap["apple"])
	delete(aMap, "banana")
	fmt.Println("Map after deleting 'banana':", aMap)

	// 4. FUNCTIONS
	// Functions are blocks of code that perform a specific task. [17]
	// They are declared with the 'func' keyword. [8]
	result := add(5, 3)
	fmt.Println("Result of add function:", result)

	// Functions can return multiple values. [17]
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	// 5. STRUCTS
	// A struct is a composite type that groups together variables under a single name.
	type Person struct {
		Name string
		Age  int
	}

	person1 := Person{Name: "Alice", Age: 30}
	fmt.Println("Person 1:", person1.Name, person1.Age)

	// You can also create a pointer to a struct.
	person2 := &Person{Name: "Bob", Age: 25}
	fmt.Println("Person 2:", person2.Name, person2.Age)

	// 6. METHODS
	// Go does not have classes, but you can define methods on types.
	// A method is a function with a special receiver argument.
	r := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of the rectangle:", r.Area())

	// 7. INTERFACES
	// Interfaces in Go provide a way to specify the behavior of an object.
	// An interface is a collection of method signatures. [17]
	// A type implements an interface by implementing its methods.
	c := Circle{Radius: 5}

	printArea(r)
	printArea(c)

	// 8. ERROR HANDLING
	// Go has a simple and explicit approach to error handling. [16]
	// Functions that can fail return an 'error' value as their last return value. [16]
	val, err := sqrt(-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", val)
	}

	// 9. GOROUTINES AND CONCURRENCY
	// Goroutines are lightweight threads managed by the Go runtime. [8]
	// You can start a new goroutine by using the 'go' keyword before a function call. [8]
	go say("world")
	say("hello")
	// Note: The program might exit before the "world" goroutine finishes.
	// We'll use channels to synchronize.

	// 10. CHANNELS
	// Channels are a typed conduit through which you can send and receive values with the 'channel operator', <-.
	// They are used to communicate between goroutines.
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks.

	// This concludes our introductory tour of Go!
	// There's much more to learn, such as the standard library, testing, and package management.
	// A great next step is to explore "A Tour of Go" and "Go by Example" for more hands-on practice. [4]
}

// A simple function that takes two integers and returns their sum.
func add(a int, b int) int {
	return a + b
}

// A function that returns multiple values.
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// A struct representing a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// A method for the Rectangle struct to calculate its area.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// A struct representing a circle.
type Circle struct {
	Radius float64
}

// A method for the Circle struct to calculate its area.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// An interface for shapes that have an Area method.
type Shape interface {
	Area() float64
}

// A function that takes a Shape interface and prints its area.
func printArea(s Shape) {
	fmt.Println("The area is:", s.Area())
}

// A function that returns an error if the input is negative.
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot take the square root of a negative number")
	}
	return math.Sqrt(x), nil
}

// A simple function for demonstrating goroutines.
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
