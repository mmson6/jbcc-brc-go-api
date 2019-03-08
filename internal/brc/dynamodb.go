package brc

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type ctxKey string

const keyDynamoDB ctxKey = "dynamo"
const dynamoURL string = "dynamodb.us-east-1.amazonaws.com"

func BuildDynamoDB(ctx context.Context) dynamodbiface.DynamoDBAPI {
	val := ctx.Value(keyDynamoDB)
	if db, ok := val.(dynamodbiface.DynamoDBAPI); ok {
		return db
	} else {
		return newDynamoDB(ctx)
	}
}

////////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

func newDynamoDB(ctx context.Context) *dynamodb.DynamoDB {
	awsSess := SharedAWSSession(ctx)
	dbCfg := aws.NewConfig().WithEndpoint(dynamoURL)
	db := dynamodb.New(awsSess, dbCfg)

	return db
}
