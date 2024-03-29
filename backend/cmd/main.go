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
	"github.com/GaryBrownEEngr/go_web_dev/backend/aws/awsDynamo"
	"github.com/GaryBrownEEngr/go_web_dev/backend/aws/awssecrets"
	"github.com/GaryBrownEEngr/go_web_dev/backend/gamestore"
	"github.com/GaryBrownEEngr/go_web_dev/backend/gcp/gcpsecretmanager"
	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/sessionuser"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
)

func run(
	getenv func(string) string,
) {
	// Setup pretend articles
	Articles := []models.Article{
		{Id: 1, Title: "Hello", Desc: "Article Description for Hello", Content: "Article Content for Hello"},
		{Id: 2, Title: "Hello2", Desc: "Article Description for Hello2", Content: "Article Content for Hello2"},
	}
	articles := articlestore.NewStore(Articles)

	cloudProvider := getenv("CLOUD_PROVIDER")

	var err error
	var secrets models.SecretStore
	var userDataDb models.KeyDBStore
	var gameDataDb models.KeyDBStore

	switch cloudProvider {
	case "AWS":
		secrets, err = awssecrets.NewAwsSecretManager(getenv)
		if err != nil {
			log.Fatal(err)
		}

		userDataDb, err = awsDynamo.NewDynamoDB(getenv, "GoWebDev_user", "username")
		if err != nil {
			log.Fatal(err)
		}

		gameDataDb, err = awsDynamo.NewDynamoDB(getenv, "GoWebDev", "Name")
		if err != nil {
			log.Fatal(err)
		}

	case "GCP":
		secrets, err = gcpsecretmanager.NewGcpSecretManager(getenv)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(secrets.Get("Best bacon"))

	// Setup AWS DynamoDB based user store
	users, err := sessionuser.NewUserStore(userDataDb)
	if err != nil {
		log.Fatal(err)
	}
	mathData, err := gamestore.NewUserStore(gameDataDb)
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

func main() {
	log.SetFlags(log.Lshortfile)
	gitHash, err := utils.GitBuildHashGet()
	if err != nil {
		log.Printf("Error while getting Git Hash: %v\n", err)
	}
	fmt.Printf("Go Web Dev: %s, %s\n", gitHash.Hash, gitHash.BuildTime)

	// Set the ENV variable "ZONEINFO" to where to find the file zoneinfo.zip.
	// This is needed by the time package to know how to interpret time zones.
	os.Setenv("ZONEINFO", "/app/zoneinfo.zip")
	// l, err := time.LoadLocation("America/Los_Angeles")

	run(os.Getenv)
}
