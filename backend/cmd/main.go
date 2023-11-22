package main

// gofumpt:diagnose version: v0.4.0 (go1.19.3) flags: -lang=v1.17 -modpath=backend
// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GaryBrownEEngr/go_web_dev/backend/api"
	"github.com/GaryBrownEEngr/go_web_dev/backend/articlestore"
	"github.com/GaryBrownEEngr/go_web_dev/backend/aws/awssecrets"
	"github.com/GaryBrownEEngr/go_web_dev/backend/gamestore"
	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/sessionuser"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
)

func main() {
	log.SetFlags(log.Lshortfile)
	gitHash, err := utils.GitBuildHashGet()
	if err != nil {
		log.Printf("Error while getting Git Hash: %v\n", err)
	}
	fmt.Printf("Go Web Dev: %s, %s\n", gitHash.Hash, gitHash.BuildTime)

	// Setup pretend articles
	Articles := []models.Article{
		{Id: 1, Title: "Hello", Desc: "Article Description for Hello", Content: "Article Content for Hello"},
		{Id: 2, Title: "Hello2", Desc: "Article Description for Hello2", Content: "Article Content for Hello2"},
	}
	articles := articlestore.NewStore(Articles)

	// Setup AWS Secret Manager
	secrets, err := awssecrets.NewSecretManager()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(secrets.Get("Best bacon"))

	// Setup AWS DynamoDB based user store
	users, err := sessionuser.NewUserStore()
	if err != nil {
		log.Fatal(err)
	}
	mathData, err := gamestore.NewUserStore()
	if err != nil {
		log.Fatal(err)
	}

	// Setup the Paseto token maker
	paseto_maker_symmetric_key, err := secrets.Get("paseto_maker_symmetric_key")
	if err != nil {
		log.Fatal(err)
	}
	tokenMaker, err := utils.NewPasetoMaker(paseto_maker_symmetric_key)
	if err != nil {
		log.Fatal(err)
	}

	// Build and run the server
	server := api.NewServer(articles, secrets, users, mathData, tokenMaker)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Setup Complete. Listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, server))
	// log.Fatal(http.ListenAndServe("localhost:10000", server))

	// go to http://localhost:10000/
	// go to http://localhost:10000/post.html
	// go to http://localhost:10000/guess.html
	// go to http://localhost:10000/test2.html
	// go to http://localhost:10000/api/articles
	// go to http://localhost:10000/ticktacktoe
}
