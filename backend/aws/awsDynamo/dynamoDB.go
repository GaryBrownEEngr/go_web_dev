package awsDynamo

import (
	"context"
	"fmt"

	"github.com/GaryBrownEEngr/twertle_api_dev/backend/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dbState struct {
	service   *dynamodb.Client
	tableName string
}

var _ models.KeyDBStore = &dbState{}

// https://dynobase.dev/dynamodb-golang-query-examples/
// https://docs.aws.amazon.com/code-library/latest/ug/go_2_dynamodb_code_examples.html
// https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/dynamodb/actions/table_basics.go#L216

func NewDynamoDB(tableName string) (*dbState, error) {
	ret := &dbState{
		tableName: tableName,
	}

	region := "us-west-2"
	// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
	config, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("Error while loading AWS default config: %w", err)
	}

	svc := dynamodb.NewFromConfig(config)
	ret.service = svc

	return ret, nil
}

func (s *dbState) Get(ctx context.Context, key string, out interface{}) error {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]types.AttributeValue{
			"Name": &types.AttributeValueMemberS{Value: key},
		},
	}
	response, err := s.service.GetItem(ctx, params)
	if err != nil {
		return fmt.Errorf("Error when getting DynamoDB Key: %s, %s: %w", s.tableName, key, err)
	}

	err = attributevalue.UnmarshalMap(response.Item, out)
	if err != nil {
		return err
	}

	return nil
}

func (s *dbState) Put(ctx context.Context, in interface{}) error {
	item, err := attributevalue.MarshalMap(in)
	if err != nil {
		return err
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String(s.tableName),
		Item:      item,
	}
	_, err = s.service.PutItem(ctx, params)
	if err != nil {
		return fmt.Errorf("Error when writing DynamoDB: %s, %#v: %w", s.tableName, in, err)
	}

	return nil
}
