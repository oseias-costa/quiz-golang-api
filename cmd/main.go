package main

import (
	"log"
	"net/http"

	"context"

	"github.com/gorilla/mux"
	"github.com/oseias-costa/quiz-golang-api/handler"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	r := mux.NewRouter()
	port := ":8000"

	opt := option.WithCredentialsFile("/home/oseias-costa/test-cfcbf-firebase-adminsdk-9keyk-2ae765abd7.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
	r.HandleFunc("/posts", handler.GetAllPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts", handler.AddPost).Methods(http.MethodPost)

	http.ListenAndServe(port, r)
}
