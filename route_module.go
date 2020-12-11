package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// 핸들러 구조체를 작성한다.
// 핸들러의 기본 내용은 경로, 핸들러함수, 메소드 이므로 3가지를 작성했다.
type handler struct {
	path    string
	fun     http.HandlerFunc
	methods []string
}

// 핸들러 구조체 목록을 받아서 핸들러를 등록한다.
func makeHandlers(hs []handler, r *mux.Router, prefix string) {
	apiRouter := r.PathPrefix(prefix).Subrouter()

	for _, h := range hs {
		if len(h.methods) == 0 {
			apiRouter.PathPrefix(h.path).HandlerFunc(h.fun)
		} else {
			apiRouter.PathPrefix(h.path).HandlerFunc(h.fun).Methods(h.methods...)
		}
	}
}
