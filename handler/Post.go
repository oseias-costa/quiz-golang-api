package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/oseias-costa/quiz-golang-api/entity"
	"github.com/oseias-costa/quiz-golang-api/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error getting posts"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "error with unmarshall"}`))
		return
	}

	post.Id = rand.Int63()
	repo.Save(&post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
