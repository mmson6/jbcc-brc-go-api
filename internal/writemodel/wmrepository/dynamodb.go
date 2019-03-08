package wmrepository

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/jbcc/brc-api/internal/models"
	"github.com/jbcc/brc-api/pkg/brcapiv1"
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
// PUBLIC FUNCTIONS

//////////
// Interface: Repository

func (repo *DynamoRepository) CreateInitialRecordForNewUser(ctx context.Context, id string, userProfile brcapiv1.UserProfile) error {
	model := models.UserRecord{
		ID:          id,
		DisplayName: userProfile.DisplayName,
		Group:       userProfile.Group,
	}

	err := repo.putUserRecord(ctx, model)
	if err != nil {
		log.Printf("module=wm.dynamodb action=putUserRecord, state=error")
		return err
	}

	return nil
}

func (repo *DynamoRepository) CreateNewUser(ctx context.Context, id string, userProfile brcapiv1.UserProfile) error {
	model := models.UserProfile{
		ID:          id,
		DisplayName: userProfile.DisplayName,
		Group:       userProfile.Group,
	}

	err := repo.putUserProfile(ctx, model)
	if err != nil {
		log.Printf("module=wm.dynamodb action=putUserProfile, state=error")
		return err
	}

	return nil
}

func (repo *DynamoRepository) UpdateUserRecord(ctx context.Context, userRecord models.UserRecord) error {
	err := repo.putUserRecord(ctx, userRecord)
	if err != nil {
		log.Printf("module=wm.dynamodb action=putUserRecord, state=error")
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func (repo *DynamoRepository) putItem(ctx context.Context, tblItem tableItem) error {
	log.Printf("module=wm.repository action=putitem, state=begin, pk=%s, sk=%s", tblItem.PartitionKey, tblItem.SortKey)

	item, err := dynamodbattribute.MarshalMap(tblItem)
	if err != nil {
		log.Printf("module=wm.repository action=putitem, state=error, pk=%s, sk=%s, error=%s", tblItem.PartitionKey, tblItem.SortKey, err.Error())
		return err
	}

	input := dynamodb.PutItemInput{
		TableName: aws.String("jbcc.brc.v1"),
		Item:      item,
	}

	_, err = repo.client.PutItemWithContext(ctx, &input)
	if err != nil {
		log.Printf("module=wm.repository action=putitem, state=error, pk=%s, sk=%s, error=%s", tblItem.PartitionKey, tblItem.SortKey, err.Error())
		return err
	}

	log.Printf("module=wm.repository action=putitem, state=success, pk=%s, sk=%s", tblItem.PartitionKey, tblItem.SortKey)

	return nil
}

//////////
// User

func (repo *DynamoRepository) putUserProfile(ctx context.Context, userProfile models.UserProfile) error {
	item := itemForUserProfile(userProfile)

	return repo.putItem(ctx, item)
}

func (repo *DynamoRepository) putUserRecord(ctx context.Context, userRecord models.UserRecord) error {
	item := itemForUserRecord(userRecord)

	return repo.putItem(ctx, item)
}
