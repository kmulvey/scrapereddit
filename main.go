package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
	var idFlag = flag.String("id", "", "oauth id")
	var secretFlag = flag.String("secret", "", "oauth secret")
	var usernameFlag = flag.String("reddit username", "", "reddit username")
	var passwordFlag = flag.String("password", "", "reddit password")
	flag.Parse()

	credentials := reddit.Credentials{ID: *idFlag, Secret: *secretFlag, Username: *usernameFlag, Password: *passwordFlag}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		log.Fatal(err)
	}
	var posts, errs, _ = reddit.DefaultClient().Stream.Posts("AVexchange", reddit.StreamMaxRequests(0))
	for {
		select {
		case post := <-posts:

			if strings.Contains(strings.ToLower(post.Title), "[wts]") {
				if strings.Contains(strings.ToLower(post.Title), "zmf ") {
					printAndSave(client, post)
				} else if strings.Contains(strings.ToLower(post.Title), "tor ") {
					printAndSave(client, post)
				} else if strings.Contains(strings.ToLower(post.Title), "rme ") {
					printAndSave(client, post)
				}
			}
		case err := <-errs:
			log.Error(err)
		}
	}
}

func printAndSave(client *reddit.Client, post *reddit.Post) {
	fmt.Printf("%+v \n %+v \n %+v\n\n", post.Title, post.Created, post.URL)
	_, err := client.Post.Save(context.Background(), post.ID)
	if err != nil {
		log.Error("error saving ", err)
	}
}
