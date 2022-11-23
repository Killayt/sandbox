package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type ViewData struct {
	Available bool
}

func main() {
	data := ViewData{
		Available: true,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("web/templates/index.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
