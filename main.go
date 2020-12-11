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
	r.HandleFunc("/", HelloWorld).Methods("GET")
	r.HandleFunc("/path/value", ResponsePath).Methods("GET")
	// r.HandleFunc("/users", CreateUser).Methods("POST")
	// r.HandleFunc("/users", UpdateUser).Methods("PUT")
	// r.HandleFunc("/users", DeleteUser).Methods("DELETE")
	// r.HandleFunc("/my", CallMy).Methods("POST", "PUT")

	mainSubRouter := r.PathPrefix("/api").Subrouter()

	userSubRouter := mainSubRouter.PathPrefix("/users").Subrouter()
	userSubRouter.HandleFunc("", CreateUser).Methods("POST")
	userSubRouter.HandleFunc("/", UpdateUser).Methods("PUT")
	userSubRouter.PathPrefix("/{id}").HandlerFunc(SelectUser).Methods("GET")
	userSubRouter.PathPrefix("/{id}").HandlerFunc(DeleteUser).Methods("DELETE")

	subjectSubRouter := mainSubRouter.PathPrefix("/subjects").Subrouter()
	subjectSubRouter.PathPrefix("/").HandlerFunc(CreateSubject).Methods("POST")
	subjectSubRouter.PathPrefix("/").HandlerFunc(UpdateSubject).Methods("PUT")
	subjectSubRouter.PathPrefix("/{id}").HandlerFunc(SelectSubject).Methods("GET")
	subjectSubRouter.PathPrefix("/{id}").HandlerFunc(DeleteSubject).Methods("DELETE")

	fmt.Println("Server start on port: ", 8080)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func ResponsePath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your request path is :%q", html.EscapeString(r.URL.Path))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUser :")
}

func SelectUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SelectUser :")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateUser :")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteUser :")
}

func CallMy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Call My Page :")
}

func CreateSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateSubject :")
}

func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateSubject :")
}

func SelectSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SelectSubject :")
}

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteSubject :")
}
