package rmrepository

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

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
