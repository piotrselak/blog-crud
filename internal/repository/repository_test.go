package repository

import (
	// "os/exec"
	"fmt"
	"testing"

	"github.com/piotrselak/blog-crud/internal/db"
)

var client = db.InitializeMongoDB()
var collection = client.Database("blog").Collection("posts")

func TestGetAllPostsWhenNoneExist(t *testing.T) {

	if posts := GetAllPosts(collection); len(posts) != 0 {
		t.Error("No posts should exist")
	}
}

func TestCreatingAPost(t *testing.T) {
	err := AddNewPost(collection, "Title", "Sample text")
	if err != nil {
		t.Error("Couldn't add a new post.")
	}
	if posts := GetAllPosts(collection); len(posts) != 1 {
		fmt.Println(posts)
		t.Error("One post should exist")
	}
}
