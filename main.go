package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	mux.HandleFunc("GET /users/{id}", GetUser)
	mux.HandleFunc("DELETE /users/{id}", DeleteUser)
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, ok := usersCache[id]
	if !ok {
		http.Error(w, "User not found.", http.StatusBadRequest)
		return
	}
	cacheMutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	name, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(name)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	_, ok := usersCache[id]
	if !ok {
		http.Error(w, "User not found.", http.StatusBadRequest)
		return
	}
	cacheMutex.RUnlock()

	w.WriteHeader(http.StatusNoContent)
	delete(usersCache, id)
}
