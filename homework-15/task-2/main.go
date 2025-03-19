package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

/*
	1. Create a URL
	2. Create a HTTP client
	3. Make a GET request
	4. Parse server response
	5. Create a Post struct
*/

/*
	Additional task

	1. Use slice instead of map
	2. Capitalize the title of each post in slice of Post
*/

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var POST_URL = "https://jsonplaceholder.typicode.com/posts/"
var postsCollection = map[int]Post{} // Post database

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func() {
			defer wg.Done()
			post, err := FetchSinglePost(i)
			if err != nil {
				fmt.Printf("Error GET post: %s", err)
				return
			}
			mx.Lock()
			postsCollection[post.Id] = post
			mx.Unlock()
		}()
	}
	wg.Wait()

	PrintPosts(postsCollection)
}

func FetchSinglePost(id int) (Post, error) {
	var post Post
	client := http.Client{}

	response, err := client.Get(fmt.Sprintf("%s%d", POST_URL, id))
	if err != nil {
		return post, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func PrintPosts(postsCollection map[int]Post) {
	for k, v := range postsCollection {
		fmt.Printf("ID: %d\nTitle: %s\n", k, v.Title)
		fmt.Println("------------------")
	}
}
