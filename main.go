package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var handlers []handler
	handlers = append(handlers, usersHandlers()...)
	handlers = append(handlers, subjectsHandlers()...)

	router := mux.NewRouter()
	makeHandlers(handlers, router, "/api")

	fmt.Println("Server start on port: ", 8080)
	log.Fatal(http.ListenAndServe(":8080", router))
}
