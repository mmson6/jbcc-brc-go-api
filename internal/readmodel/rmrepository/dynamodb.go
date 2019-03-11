package rmrepository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/jbcc/brc-api/internal/models"
)

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

// Logging state values
const (
	stateError   = "error"
	stateStart   = "start"
	stateSuccess = "success"
)

////////////////////////////////////////////////////////////////////////////////
// STRUCTURES

type DynamoRepository struct {
	client dynamodbiface.DynamoDBAPI
}

////////////////////////////////////////////////////////////////////////////////
// INITIALIZERS

func NewDynamoRepository(client dynamodbiface.DynamoDBAPI) *DynamoRepository {
	return &DynamoRepository{
		client: client,
	}
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

//////////
// Interface: Repository

func (repo *DynamoRepository) CheckForUniqueDisplayName(ctx context.Context, displayName string) (bool, error) {
	item, err := repo.getItem(ctx, displayName, sortKeyUserProfile)

	if err != nil {
		return false, err
	} else if item == nil {
		return true, nil
	}

	return false, nil
}

func (repo *DynamoRepository) CheckForUniqueID(ctx context.Context, id string) (bool, error) {
	item, err := repo.getItem(ctx, id, sortKeyUserRecord)

	if err != nil {
		return false, err
	} else if item == nil {
		return true, nil
	}

	return false, nil
}

func (repo *DynamoRepository) ReadLeaderboard(ctx context.Context) (*models.Leaderboard, error) {
	items, err := repo.queryGSI1(ctx, sortKeyUserRecord)
	if err != nil {
		return nil, err
	}

	leaderboard := leaderboardForItems(items)
	return leaderboard, nil
}

func (repo *DynamoRepository) ReadUserRecordByUserID(ctx context.Context, userID string) (*models.UserRecord, error) {
	item, err := repo.getItem(ctx, userID, sortKeyUserRecord)
	if err != nil {
		return nil, err
	} else if item == nil {
		uErr := errors.New("unable to find user record")
		return nil, uErr
	}

	// Extract the model
	userRecord := userRecordForItem(*item)

	return userRecord, nil
}

func (repo *DynamoRepository) getItem(ctx context.Context, pk, sk string) (*tableItem, error) {
	log.Printf("module=rm.repository action=getitem, state=begin, pk=%s, sk=%s", pk, sk)

	input := dynamodb.GetItemInput{
		TableName: aws.String("jbcc.brc.v1"),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(pk)},
			"sk": {S: aws.String(sk)},
		},
	}

	output, err := repo.client.GetItemWithContext(ctx, &input)
	if err != nil {
		log.Printf("module=rm.repository action=getitem, state=error, pk=%s, sk=%s, error=%s", pk, sk, err.Error())
		return nil, err
	} else if len(output.Item) == 0 {
		// Not found
		log.Printf("module=rm.repository action=getitem, state=not-found, pk=%s, sk=%s", pk, sk)
		return nil, nil
	}

	var tblItem tableItem
	err = dynamodbattribute.UnmarshalMap(output.Item, &tblItem)
	if err != nil {
		log.Printf("module=rm.repository action=getitem, state=error, pk=%s, sk=%s, error=%s", pk, sk, err.Error())
		return nil, err
	}

	log.Printf("module=rm.repository action=getitem, state=success, pk=%s, sk=%s", pk, sk)

	return &tblItem, nil
}

// Find all inverse relationships for an entity.
func (repo *DynamoRepository) queryGSI1(ctx context.Context, sk string) ([]tableItem, error) {
	log.Printf("module=wm.repository action=query, target=gsi1, state=begin, pk=%s", sk)

	var err error
	var count int

	idx := "jbcc.brc.v1" + ".gsi2"

	// Build query expressions
	keyCond := expression.Key("sk").Equal(expression.Value(sk))
	proj := expression.NamesList( // Only the index and primary keys are projected into the index
		expression.Name("pk"),
		expression.Name("sk"),
		expression.Name("unique_key"),
		expression.Name("verses"),
	)
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj).
		Build()
	if err != nil {
		log.Printf("module=wm.repository action=query, target=gsi1, state=error, pk=%s, error=%s", sk, err.Error())
		return nil, err
	}

	// Create the query input
	input := dynamodb.QueryInput{
		TableName:                 aws.String("jbcc.brc.v1"),
		IndexName:                 aws.String(idx),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ScanIndexForward:          aws.Bool(false),
	}

	// Execute the query
	output, err := repo.client.QueryWithContext(ctx, &input)
	if err != nil {
		log.Printf("module=wm.repository action=query, target=gsi1, state=error, pk=%s, error=%s", sk, err.Error())
		return nil, err
	}

	count = len(output.Items)
	tblItems := make([]tableItem, 0, count)
	for _, outputItem := range output.Items {
		var tblItem tableItem
		err = dynamodbattribute.UnmarshalMap(outputItem, &tblItem)
		if err != nil {
			log.Printf("module=wm.repository action=query, target=gsi1, state=error, pk=%s, error=%s", sk, err.Error())
			return nil, err
		}
		tblItems = append(tblItems, tblItem)
	}

	return tblItems, nil
}
