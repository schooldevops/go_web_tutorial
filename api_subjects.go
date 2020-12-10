package main

import (
	"fmt"
	"net/http"
)

func subjectsHandlers() []handler {
	return []handler{
		{path: "/subjects/{id}", fun: SubjectByID, methods: []string{"GET"}},
		{path: "/subjects/{id}", fun: DeleteSubjectByID, methods: []string{"DELETE"}},
		{path: "/subjects", fun: CreateSubjects, methods: []string{"POST"}},
		{path: "/subjects", fun: AllSubjects, methods: []string{"GET"}},
		{path: "/subjects", fun: UpdateSubjects, methods: []string{"PUT"}},
	}
}

func CreateSubjects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateSubjects")
}

func AllSubjects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AllSubjects")
}

func SubjectByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SubjectByID")
}

func UpdateSubjects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateSubjects")
}

func DeleteSubjectByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteSubjectByID")
}
