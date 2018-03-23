package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("users list"))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("users 2"))
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("users 3"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
