package repository

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/oseias-costa/quiz-golang-api/entity"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "test-cfcbf"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":   post.Id,
		"Name": post.Name,
		"Age":  post.Age,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post

	fmt.Println("All posts:")
	iter := client.Collection(collectionName).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("error to iterate on collection %s: %v", collectionName, err)
		}

		post := entity.Post{
			Id:   doc.Data()["Id"].(int64),
			Name: doc.Data()["Name"].(string),
			Age:  doc.Data()["Age"].(int64),
		}
		posts = append(posts, post)
		fmt.Println(doc.Data())
	}

	return posts, nil
}
