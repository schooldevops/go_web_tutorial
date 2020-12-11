package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	http.HandleFunc("/path/value", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Your request path is :%q", html.EscapeString(r.URL.Path))
	})

	fmt.Println("Server start on port: ", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
