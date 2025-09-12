// Filename: concurrency.go
//
// This Go file is a deep dive into the concurrency model of the Go programming language.
// It is crafted for developers who have a basic understanding of Go and wish to master its powerful concurrency primitives.
// We will explore goroutines, channels, and synchronization techniques not just as language features,
// but as tools for building robust, efficient, and scalable concurrent applications.
//
// The file is structured as a tutorial, with each section building upon the last.
// You can read and run this file from top to bottom to follow the learning path.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SECTION 1: GOROUTINES - THE HEARTBEAT OF GO CONCURRENCY
//
// Go's approach to concurrency is centered around the concept of a 'goroutine'.
// A goroutine is a lightweight thread of execution managed by the Go runtime.
// Unlike traditional OS threads, which can be resource-intensive, goroutines are cheap.
// They start with a small stack size (a few kilobytes) that can grow and shrink as needed.
// It is common and practical to have hundreds of thousands, or even millions, of goroutines running in a single program.

// We start a goroutine by using the 'go' keyword followed by a function call.
func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		// We use time.Sleep to simulate work and to allow the Go scheduler
		// to switch between goroutines, demonstrating their concurrent execution.
		time.Sleep(time.Millisecond * 120)
	}
}

// The 'sync.WaitGroup' is a crucial tool for synchronizing goroutines.
// It provides a mechanism to wait for a collection of goroutines to finish their execution.
// The main goroutine can block until all other goroutines in the WaitGroup have completed.
func useWaitGroup() {
	// A WaitGroup is a counter that blocks until it reaches zero.
	var wg sync.WaitGroup

	fmt.Println("WaitGroup Demo: Starting goroutines...")

	// We have 2 goroutines to wait for, so we call Add(2).
	wg.Add(2)

	// We start our goroutines. It's common practice to pass the WaitGroup
	// pointer to the function that the goroutine will execute.
	go func() {
		// defer wg.Done() is a robust way to ensure that the counter is
		// decremented when the function exits, regardless of how it exits.
		defer wg.Done()
		say("Hello")
	}()

	go func() {
		defer wg.Done()
		say("World")
	}()

	// wg.Wait() blocks the execution of this function until the WaitGroup counter becomes 0.
	// This happens after both goroutines have called wg.Done().
	fmt.Println("WaitGroup Demo: Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("WaitGroup Demo: Goroutines have finished.")
}

// SECTION 2: CHANNELS - THE ARTERIES OF COMMUNICATION
//
// Channels are the pipes that connect concurrent goroutines. You can send values into channels
// from one goroutine and receive those values into another goroutine.
// "Do not communicate by sharing memory; instead, share memory by communicating." - Effective Go

// --- Unbuffered Channels ---
// By default, channels are unbuffered. This means they will only accept a send (chan <-)
// if there is a corresponding receive (<- chan) ready to take the value.
// They are powerful synchronization tools because they force a rendezvous between the sender and receiver.
func unbufferedChannelDemo() {
	fmt.Println("\nUnbuffered Channel Demo: Starting...")

	// Create an unbuffered channel of strings.
	messages := make(chan string)

	go func() {
		fmt.Println("Unbuffered: Goroutine is sending 'ping'...")
		// This send operation will BLOCK until another goroutine is ready to receive.
		messages <- "ping"
		fmt.Println("Unbuffered: Goroutine sent 'ping'.")
	}()

	// Sleep for a moment to demonstrate that the sending goroutine is blocked.
	time.Sleep(time.Second * 1)
	fmt.Println("Unbuffered: Main is ready to receive.")

	// This receive operation will block until a value is sent on the channel.
	msg := <-messages
	fmt.Println("Unbuffered: Main received message:", msg)
}

// --- Buffered Channels ---
// Buffered channels accept a limited number of values without a corresponding receiver.
// They are useful when you want to decouple the sender and receiver, or when you have
// a burst of sends that you want to handle asynchronously.
func bufferedChannelDemo() {
	fmt.Println("\nBuffered Channel Demo: Starting...")
	// Create a buffered channel of strings with a capacity of 2.
	messages := make(chan string, 2)

	// Because the channel is buffered, we can send these values without a
	// corresponding receiver ready. The send operations do not block.
	messages <- "buffered"
	messages <- "channel"
	fmt.Println("Buffered: Sent two messages without blocking.")

	// If we try to send a third value, the send will block until a value is received.
	// go func() { messages <- "overflow" }() // Uncommenting this would block.

	// We can receive the values that were buffered.
	fmt.Println("Buffered: Received:", <-messages)
	fmt.Println("Buffered: Received:", <-messages)
}

// --- The 'select' Statement ---
// The 'select' statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case.
// It chooses one at random if multiple are ready.
func selectDemo() {
	fmt.Println("\nSelect Statement Demo: Starting...")
	c1 := make(chan string)
	c2 := make(chan string)

	// These goroutines will send messages on their respective channels after a delay.
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// The select statement will wait for messages from both channels.
	// It will run the case for the first channel that becomes ready.
	for i := 0; i < 2; i++ {
		fmt.Println("Select: Waiting for a message...")
		select {
		case msg1 := <-c1:
			fmt.Println("Select: Received from c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Select: Received from c2:", msg2)
		case <-time.After(time.Second * 3):
			// time.After is a common way to implement timeouts in a select.
			// It returns a channel that receives a value after the specified duration.
			fmt.Println("Select: Timeout waiting for message.")
		}
	}
}

// SECTION 3: COMMON CONCURRENCY PATTERNS

// --- Worker Pool (Fan-out, Fan-in) ---
// This is a very common and powerful pattern. We have a set of jobs to do,
// a number of workers to do them, and we need to collect the results.
// - Fan-out: A single source distributes jobs to multiple workers.
// - Fan-in: Multiple sources consolidate results into a single channel.

// The worker function defines the work to be done.
// It receives jobs from a 'jobs' channel and sends results to a 'results' channel.
func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		// Simulate work
		time.Sleep(time.Second)
		// Send the result of the work to the results channel.
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func workerPoolDemo() {
	fmt.Println("\nWorker Pool Demo: Starting...")
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// FAN-OUT: Start the workers.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, &wg, jobs, results)
	}

	// Send jobs to the workers.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Close the jobs channel to signal to the workers that no more jobs will be sent.
	close(jobs)

	// FAN-IN: Wait for all workers to finish their jobs.
	// We must do this in a separate goroutine to avoid a deadlock.
	// If we waited here, the results channel would never be read from.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect the results.
	// This loop will block until the 'results' channel is closed.
	for r := range results {
		fmt.Println("Collected result:", r)
	}
	fmt.Println("Worker Pool Demo: All jobs completed.")
}

// SECTION 4: HANDLING CONCURRENCY HAZARDS

// --- Race Conditions and Mutexes ---
// A race condition occurs when multiple goroutines access shared data concurrently,
// and at least one of them modifies it. The result of the execution depends on the
// non-deterministic ordering of operations.

// We can detect race conditions using the Go race detector: `go run -race your_file.go`

// To prevent race conditions, we use synchronization primitives. The most common is the `sync.Mutex`.
// A Mutex (Mutual Exclusion lock) ensures that only one goroutine can access a critical section of code at a time.
type SafeCounter struct {
	mu sync.Mutex // A Mutex to protect the counter
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()   // Lock the mutex before accessing the map.
	defer c.mu.Unlock() // Defer the Unlock to ensure it happens when the function exits.
	c.v[key]++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func mutexDemo() {
	fmt.Println("\nMutex Demo: Starting...")
	sc := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup
	numIncrements := 1000

	wg.Add(numIncrements)
	// Start 1000 goroutines that all try to increment the same counter.
	// Without a mutex, this would be a classic race condition.
	for i := 0; i < numIncrements; i++ {
		go func() {
			defer wg.Done()
			sc.Inc("somekey")
		}()
	}

	wg.Wait()
	fmt.Printf("Mutex Demo: Final counter value should be %d, and it is: %d\n", numIncrements, sc.Value("somekey"))
}

// SECTION 5: THE 'CONTEXT' PACKAGE - MANAGING GOROUTINE LIFECYCLES

// The 'context' package is essential for managing the lifecycle of goroutines,
// especially in server applications. It allows you to pass cancellation signals,
// deadlines, and other request-scoped values across API boundaries to all the
// goroutines involved in handling a request.

// This function simulates a long-running operation that can be cancelled.
func longOperation(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Context: Long operation started.")
	select {
	case <-time.After(time.Second * 5):
		// The operation completed successfully.
		fmt.Println("Context: Long operation finished successfully.")
	case <-ctx.Done():
		// The context was cancelled. ctx.Err() provides the reason for cancellation.
		fmt.Println("Context: Long operation cancelled:", ctx.Err())
	}
}

func contextDemo() {
	fmt.Println("\nContext Demo: Starting...")
	var wg sync.WaitGroup
	wg.Add(1)

	// context.WithTimeout returns a new context that will be automatically cancelled
	// after the specified duration.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	// It's good practice to call the cancel function, even if the context times out.
	// This releases any resources associated with the context.
	defer cancel()

	go longOperation(ctx, &wg)

	wg.Wait()
	fmt.Println("Context Demo: Finished.")
}

// main is the entry point of our program.
func main() {
	// Execute the demonstrations in order.
	fmt.Println("--- START: SECTION 1: GOROUTINES ---")
	useWaitGroup()
	fmt.Println("--- END: SECTION 1 ---")

	fmt.Println("\n--- START: SECTION 2: CHANNELS ---")
	unbufferedChannelDemo()
	bufferedChannelDemo()
	selectDemo()
	fmt.Println("--- END: SECTION 2 ---")

	fmt.Println("\n--- START: SECTION 3: CONCURRENCY PATTERNS ---")
	workerPoolDemo()
	fmt.Println("--- END: SECTION 3 ---")

	fmt.Println("\n--- START: SECTION 4: CONCURRENCY HAZARDS ---")
	mutexDemo()
	fmt.Println("--- END: SECTION 4 ---")

	fmt.Println("\n--- START: SECTION 5: THE CONTEXT PACKAGE ---")
	contextDemo()
	fmt.Println("--- END: SECTION 5 ---")

	fmt.Println("\nConcurrency in Go tutorial completed.")
}
