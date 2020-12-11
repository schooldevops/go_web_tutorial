package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func usersHandlers() []handler {
	return []handler{
		{path: "/users/{id}/summary/{key}", fun: GetUserSummaryByKey, methods: []string{"GET"}},
		{path: "/users/{id:[0-9]+}", fun: UserByID, methods: []string{"GET"}},
		{path: "/users/secret", fun: NotSupportedFunction, methods: []string{"GET"}},
		{path: "/users/{id}", fun: UserByID, methods: []string{"GET"}},
		{path: "/users/{id}", fun: DeleteUserByID, methods: []string{"DELETE"}},
		{path: "/users", fun: CreateUser, methods: []string{"POST"}},
		{path: "/users", fun: AllUsers, methods: []string{"GET"}},
		{path: "/users", fun: UpdateUser, methods: []string{"PUT"}},
		{path: "/users/secret", fun: NotSupportedFunction, methods: []string{"GET"}},
	}
}

func NotSupportedFunction(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusNotImplemented, "Your request not implemented yet.")
}

func GetUserSummaryByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	summaryKey := vars["key"]

	// fmt.Fprintf(w, "UserByID :%v, key: %v", userID, summaryKey)
	payload := map[string]string{"message": fmt.Sprintf("UserByID :%v, key: %v", userID, summaryKey)}
	responseWithJSON(w, http.StatusOK, payload)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUser")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()

	page := values["page"]
	rows := values["rows"]
	startName := values["startName"]

	fmt.Println("Hostname: ", r.URL.Hostname())

	fmt.Fprintf(w, "AllUsers: %v %v %v", page, rows, startName)
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
