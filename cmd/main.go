package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/static/about.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/static/index.html")
	})

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func newFunction() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/", fs)
}
