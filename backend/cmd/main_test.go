package main

import (
	"log"
	"testing"
)

func Test_run(t *testing.T) {
	t.SkipNow()

	getenv := func(in string) string {
		switch in {
		case "CLOUD_PROVIDER":
			return "AWS"
		case "PORT":
			return "8080"
		case "AWS_SECRET_NAME":
			return "AppRunner/GoWebDev"
		case "AWS_REGION":
			return "us-west-2"
		default:
			log.Println("unknown env variable: ", in)
			return ""
		}
	}

	run(getenv)
}
