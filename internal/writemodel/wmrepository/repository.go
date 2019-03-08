package wmrepository

import (
	"context"

	"github.com/jbcc/brc-api/internal/models"
	"github.com/jbcc/brc-api/pkg/brcapiv1"
)

type Repository interface {
	// Create initial record for a new user
	CreateInitialRecordForNewUser(ctx context.Context, id string, userProfile brcapiv1.UserProfile) error

	// Create a new user
	CreateNewUser(ctx context.Context, id string, userProfile brcapiv1.UserProfile) error

	UpdateUserRecord(ctx context.Context, userRecord models.UserRecord) error
}
