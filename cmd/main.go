package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func boot() {
	fmt.Println("Server is listening...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(strings.ToUpper("Error with http.ListenAndServe\n\n"), err)
	}

}

func main() {
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/contact.html")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/about.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})

	boot()
}
