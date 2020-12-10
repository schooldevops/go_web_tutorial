package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func usersHandlers() []handler {
	return []handler{
		{path: "/users/{id}/summary/{key}", fun: GetUserSummaryByKey, methods: []string{"GET"}},
		{path: "/users/{id}", fun: UserByID, methods: []string{"GET"}},
		{path: "/users/{id}", fun: DeleteUserByID, methods: []string{"DELETE"}},
		{path: "/users", fun: CreateUser, methods: []string{"POST"}},
		{path: "/users", fun: AllUsers, methods: []string{"GET"}},
		{path: "/users", fun: UpdateUser, methods: []string{"PUT"}},
	}
}

func GetUserSummaryByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	summaryKey := vars["key"]

	fmt.Fprintf(w, "UserByID :%v, key: %v", userID, summaryKey)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUser")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AllUsers")
}

func UserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "UserByID :%v", userID)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateUser")
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteUserByID")
}
