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
		fmt.Fprint(w, "This is CONTACT page")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from ABOUT")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/tema.html")
	})

	boot()
}
