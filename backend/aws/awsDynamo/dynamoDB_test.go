package awsDynamo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDynamoDB(t *testing.T) {
	t.SkipNow()
	got, err := NewDynamoDB("GoWebDev")
	require.NoError(t, err)

	type t1 struct {
		Name      string
		FirstName string `dynamodbav:"first_name"`
		Age       int
	}

	var d1 t1
	err = got.Get(context.Background(), "Gary", &d1)
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
}
