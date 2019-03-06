package rmrepository

import (
	"context"

	"github.com/jbcc/brc-api/pkg/brcapiv1"
)

type Repository interface {
	// Create or update user identity
	UpsertUserIdentity(ctx context.Context, identity brcapiv1.UserIdentity) error
}