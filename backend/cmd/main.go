package main

// gofumpt:diagnose version: v0.4.0 (go1.19.3) flags: -lang=v1.17 -modpath=backend
// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GaryBrownEEngr/twertle_api_dev/backend/api"
	"github.com/GaryBrownEEngr/twertle_api_dev/backend/articlestore"
	"github.com/GaryBrownEEngr/twertle_api_dev/backend/aws/awsDynamo"
	"github.com/GaryBrownEEngr/twertle_api_dev/backend/aws/awssecrets"
	"github.com/GaryBrownEEngr/twertle_api_dev/backend/models"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles := []models.Article{
		{Id: 1, Title: "Hello", Desc: "Article Description for Hello", Content: "Article Content for Hello"},
		{Id: 2, Title: "Hello2", Desc: "Article Description for Hello2", Content: "Article Content for Hello2"},
	}

	articles := articlestore.NewStore(Articles)

	secrets, err := awssecrets.NewSecretManager()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(secrets.Get("Best bacon"))

	dynamodbHandle, err := awsDynamo.NewDynamoDB("GoWebDev")
	if err != nil {
		log.Fatal(err)
	}
	type t2 struct {
		Name     string
		Problems []string
	}

	var d2 t2
	err = dynamodbHandle.Get(context.TODO(), "Bacon", &d2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d2)

	server := api.NewServer(articles, secrets, dynamodbHandle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, server))
	// log.Fatal(http.ListenAndServe("localhost:10000", server))

	// go to http://localhost:10000/
	// go to http://localhost:10000/post.html
	// go to http://localhost:10000/guess.html
	// go to http://localhost:10000/test2.html
	// go to http://localhost:10000/api/articles
	// go to http://localhost:10000/ticktacktoe
}
