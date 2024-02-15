package awsDynamo

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewDynamoDB(t *testing.T) {
	t.SkipNow()

	getenv := func(in string) string {
		switch in {
		case "AWS_REGION":
			return "us-west-2"
		default:
			log.Println("unknown env variable: ", in)
			return ""
		}
	}

	got, err := NewDynamoDB(getenv, "GoWebDev", "Name")
	require.NoError(t, err)

	type t1 struct {
		Name      string
		FirstName string `dynamodbav:"first_name"`
		Age       int
	}

	var d1 t1
	err = got.Get(context.Background(), "Gary", &d1)
	require.NoError(t, err)

	// What happens when the key isn't there?
	d1 = t1{}
	err = got.Get(context.Background(), "2355634756735664", &d1)
	require.NoError(t, err)

	type t2 struct {
		Name     string
		Problems []string
	}

	var d2 t2
	err = got.Get(context.Background(), "Bacon", &d2)
	require.NoError(t, err)

	d2.Problems = append(d2.Problems, "Bacon+Turkey=food")
	err = got.Put(context.Background(), &d2)
	require.NoError(t, err)

	//
	//
	got, err = NewDynamoDB(getenv, "GoWebDev_user", "username")
	require.NoError(t, err)
	type DbUser struct {
		Username       string    `dynamodbav:"username"`
		HashedPassword string    `dynamodbav:"hashed_password"`
		CreatedAt      time.Time `dynamodbav:"created_at"`
	}
	var d3 DbUser
	err = got.Get(context.Background(), "Bacon", &d3)
	require.NoError(t, err)
}
