package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oseias-costa/quiz-golang-api/handler"
)

func main() {
	r := mux.NewRouter()
	port := ":8000"

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
	r.HandleFunc("/posts", handler.GetAllPosts).Methods(http.MethodGet)

	http.ListenAndServe(port, r)
}
