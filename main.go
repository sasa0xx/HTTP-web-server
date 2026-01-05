package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleRoot)
	http.ListenAndServe(":8080", mux)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World! :D"))
}
