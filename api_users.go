package main

import (
	"fmt"
	"net/http"
)

func usersHandlers() []handler {
	return []handler{
		{path: "/users/{id}", fun: UserByID, methods: []string{"GET"}},
		{path: "/users/{id}", fun: DeleteUserByID, methods: []string{"DELETE"}},
		{path: "/users", fun: CreateUser, methods: []string{"POST"}},
		{path: "/users", fun: AllUsers, methods: []string{"GET"}},
		{path: "/users", fun: UpdateUser, methods: []string{"PUT"}},
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUser")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AllUsers")
}

func UserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserByID")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateUser")
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteUserByID")
}
