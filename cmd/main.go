package main

import (
	"log"
	"net/http"

	"context"

	"github.com/oseias-costa/quiz-golang-api/controller"
	router "github.com/oseias-costa/quiz-golang-api/http"
	"github.com/oseias-costa/quiz-golang-api/repository"
	"github.com/oseias-costa/quiz-golang-api/service"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	postRepository repository.PostRepository = repository.NewPostRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
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

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
	httpRouter.GET("/posts", postController.GetAllPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
