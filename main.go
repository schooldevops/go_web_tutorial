package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/path/value", ResponsePath)
	http.Handle("/", r)

	fmt.Println("Server start on port: ", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func ResponsePath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your request path is :%q", html.EscapeString(r.URL.Path))
}
