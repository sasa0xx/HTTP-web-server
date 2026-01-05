package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var usersCache = make(map[int]User)
var cacheMutex sync.RWMutex

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleRoot)
	mux.HandleFunc("POST /users", CreateUser)
	http.ListenAndServe(":8080", mux)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World! :D"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(w, "User name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	usersCache[len(usersCache)] = user
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)
}
