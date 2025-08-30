# Iris
Go Learning Repository

[![Go Version](https://img.shields.io/badge/go-1.25%2B-blue.svg)](https://go.dev/doc/install/)

## Checklist
Track your progress by checking off items as you learn, practice, and implement them in projects.

---

### **Phase 1: Go Fundamentals**

*   [ ] **Setup & Tooling**
    *   [ ] Install Go
    *   [ ] Configure your `$GOPATH` and `$GOROOT`
    *   [ ] Set up your code editor (e.g., VS Code with Go extension)
    *   [ ] Write and run "Hello, World!"
    *   [ ] Learn `go run`, `go build`, `go fmt`, `go vet`
*   [ ] **Basic Syntax & Concepts**
    *   [ ] Packages and Imports
    *   [ ] Variables (`var`, `:=`) and Constants (`const`)
    *   [ ] Basic Data Types: `int`, `float64`, `string`, `bool`
    *   [ ] Pointers
    *   [ ] Functions (declaration, parameters, multiple return values)
*   [ ] **Control Flow**
    *   [ ] `if / else` statements
    *   [ ] `for` loops (all variations)
    *   [ ] `switch` statements
    *   [ ] `defer`, `panic`, and `recover`
*   [ ] **Core Data Structures**
    *   [ ] Arrays
    *   [ ] Slices (and slice operations like `append`, `make`)
    *   [ ] Maps
    *   [ ] Structs

#### **Phase 1 Projects:**
*   [ ] **CLI Number Guessing Game:** A simple game where the computer picks a number and the user has to guess it.
*   [ ] **Simple Calculator:** A command-line tool that takes two numbers and an operator to perform a calculation.
*   [ ] **Word/Character Counter:** A tool that reads a string or a file and counts the frequency of each word or character.

---

### **Phase 2: Intermediate Go & Software Structuring**

*   [ ] **Organizing Code**
    *   [ ] Creating your own packages
    *   [ ] Understanding package visibility (exported vs. unexported names)
    *   [ ] Go Modules (`go mod init`, `go get`, `go mod tidy`)
*   [ ] **Advanced Types & Methods**
    *   [ ] Methods (functions with a receiver)
    *   [ ] Interfaces (and satisfying an interface implicitly)
    *   [ ] The `error` type and idiomatic error handling
*   [ ] **Working with Data**
    *   [ ] Reading and writing files (`io`, `os` packages)
    *   [ ] Handling JSON (`encoding/json` package)

#### **Phase 2 Projects:**
*   [ ] **CLI To-Do List:** An application to add, list, and complete tasks, storing the data in a JSON file.
*   [ ] **Configuration File Parser:** A package that can read a configuration file (e.g., JSON, YAML) into a Go struct.

---

### **Phase 3: Concurrency**

*   [ ] **Core Concurrency Primitives**
    *   [ ] Goroutines
    *   [ ] Channels (buffered and unbuffered)
    *   [ ] `select` statement for managing multiple channels
*   [ ] **Concurrency Patterns**
    *   [ ] Worker pools
    *   [ ] Fan-in, Fan-out
*   [ ] **Synchronization**
    *   [ ] Mutex (`sync.Mutex`, `sync.RWMutex`)
    *   [ ] WaitGroups (`sync.WaitGroup`)
    *   [ ] Race condition detection (`go run -race`)

#### **Phase 3 Projects:**
*   [ ] **Concurrent Web Scraper:** A tool that fetches multiple URLs concurrently to speed up data gathering.
*   [ ] **Parallel File Processor:** A program that processes multiple files in a directory concurrently (e.g., resizing images, parsing logs).

---

### **Phase 4: Web Development & APIs**

*   [ ] **Standard Library for Web**
    *   [ ] Understanding the `net/http` package
    *   [ ] Building a basic web server
    *   [ ] Handling HTTP requests (GET, POST, etc.)
    *   [ ] Routing requests
*   [ ] **Building APIs**
    *   [ ] Designing RESTful API endpoints
    *   [ ] Returning JSON responses
    *   [ ] Parsing JSON from request bodies
*   [ ] **Working with a Database**
    *   [ ] Using the `database/sql` package
    *   [ ] Connecting to a SQL database (e.g., PostgreSQL, SQLite)
    *   [ ] Performing CRUD (Create, Read, Update, Delete) operations
*   [ ] **Exploring Web Frameworks (Optional)**
    *   [ ] Research popular frameworks like Gin, Echo, or Chi
    *   [ ] Rebuild an API project using a framework

#### **Phase 4 Projects:**
*   [ ] **Simple Blog API:** A REST API with endpoints for creating, reading, updating, and deleting blog posts.
*   [ ] **URL Shortener Service:** A web service that takes a long URL and returns a shortened one that redirects to the original.

---

### **Phase 5: Testing, Tooling, & Advanced Topics**

*   [ ] **Testing**
    *   [ ] Unit testing with the `testing` package
    *   [ ] Writing benchmarks
    *   [ ] Table-driven tests
*   [ ] **Advanced Concepts**
    *   [ ] The `context` package for cancellation and deadlines
    *   [ ] Reflection (`reflect` package)
    *   [ ] `cgo` for calling C code (awareness is enough)
*   [ ] **Deployment & Ecosystem**
    *   [ ] Cross-compilation
    *   [ ] Building a Docker container for a Go application
    *   [ ] Explore gRPC as an alternative to REST

#### **Phase 5 Projects:**
*   [ ] **Add Comprehensive Tests:** Go back to previous projects and add a robust suite of unit tests.
*   [ ] **Microservice:** Create a small, single-purpose service (e.g., an authentication service) that communicates over gRPC or REST.
*   [ ] **Contribute to an Open Source Project:** Find a beginner-friendly issue on a Go project and submit a pull request.
