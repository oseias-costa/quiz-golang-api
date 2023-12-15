package handler

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var posts []Post

func init() {
	posts = []Post{{Id: 1, Name: "Oséias Costa", Age: 32}, {Id: 2, Name: "Leo Borilli", Age: 29}}
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error json"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "error with unmarshall"}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	newPost, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "errpr with unmarshall"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(newPost)
}
