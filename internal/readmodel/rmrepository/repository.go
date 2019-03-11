package rmrepository

import (
	"context"

	// "github.com/jbcc/brc-api/pkg/brcapiv1"
	"github.com/jbcc/brc-api/internal/models"
)

type Repository interface {
	// Check if duplicate display name already exists
	CheckForUniqueDisplayName(ctx context.Context, displayName string) (bool, error)

	// CHeck if duplicate user ID already exists
	CheckForUniqueID(ctx context.Context, displayName string) (bool, error)

	// Read current leaderboard of BRC
	ReadLeaderboard(ctx context.Context) (*models.Leaderboard, error)

	// Read user record data
	ReadUserRecordByUserID(ctx context.Context, userID string) (*models.UserRecord, error)
}
