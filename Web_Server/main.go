package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID   int
	Name string `json:"name"`
	Age  int
}

var userCache = make(map[int]User)

var CacheMutex sync.RWMutex

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root!")
}


func deleteUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(
			w, 
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(
			w, "User not found", http.StatusNotFound,
		)
		return
	}
	CacheMutex.Lock()
	delete(userCache, id)
	defer CacheMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))	
	if err != nil{
		http.Error(
			w, 
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	CacheMutex.RLock()
	user, ok := userCache[id]
	CacheMutex.RUnlock()
	if !ok{
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil{
		http.Error(
			w,
			"User not found",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == 0 || user.Name == "" || user.Age == 0 {
		http.Error(w, "Invalid/missing required fields", http.StatusBadRequest)
		return
	}

	// Add user to cache
	CacheMutex.Lock()
	userCache[len(userCache)+1] = user
	CacheMutex.Unlock()

	// 204 No Content
	w.WriteHeader(http.StatusNoContent)
	userCache[user.ID] = user
	fmt.Fprintf(w, "User created!")
}