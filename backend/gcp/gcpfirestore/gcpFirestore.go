package gcpfirestore

// import (
// 	"context"
// 	"fmt"

// 	"cloud.google.com/go/firestore"
// 	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type dbState struct {
// 	client         *firestore.Client
// 	collection     *firestore.CollectionRef
// 	collectionName string
// }

// var _ models.KeyDBStore = &dbState{}

// // https://dynobase.dev/dynamodb-golang-query-examples/
// // https://docs.aws.amazon.com/code-library/latest/ug/go_2_dynamodb_code_examples.html
// // https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/dynamodb/actions/table_basics.go#L216

// func New(getenv func(string) string, collectionName string) (*dbState, error) {
// 	projectId := getenv("GCP_PROJECT_ID")

// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, projectId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	collection := client.Collection(collectionName)

// 	ret := &dbState{
// 		collectionName: collectionName,
// 		client:         client,
// 		collection:     collection,
// 	}

// 	return ret, nil
// }

// func (s *dbState) Get(ctx context.Context, key string, out interface{}) error {
// 	item := s.collection.Doc(key)
// 	docSnap, err := item.Get(ctx)
// 	if err != nil {
// 		if status.Code(err) == codes.NotFound {
// 			return models.KeyDbNoDocFound
// 		}
// 		return err
// 	}

// 	err = docSnap.DataTo(out)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func (s *dbState) Put(ctx context.Context, in interface{}) error {

// 	co := s.collection.Doc("Colorado")
// 	wr, err := co.Update(ctx, []firestore.Update{
// 		{Path: "cities", Value: firestore.ArrayUnion("Broomfield")},
// 	})
// 	if err != nil {
// 		// TODO: Handle error.
// 	}

// 	return nil
// }
// func (s *dbState) Delete(ctx context.Context, key string) error {
// 	return nil
// }

// func update() {
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, "project-id")
// 	if err != nil {
// 		// TODO: Handle error.
// 	}
// 	defer client.Close()

// 	fmt.Println(wr.UpdateTime)
// }
