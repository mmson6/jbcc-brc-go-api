package wmrepository

import (
	"context"

	"github.com/jbcc/brc-api/internal/brc"
)

type ctxKey string

const keyRepo ctxKey = "wm.repo"

func Build(ctx context.Context) Repository {
	val := ctx.Value(keyRepo)
	if repo, ok := val.(Repository); ok {
		return repo
	} else {
		db := brc.BuildDynamoDB(ctx)
		return NewDynamoRepository(db)
	}
}

func Inject(ctx context.Context, repo Repository) context.Context {
	return context.WithValue(ctx, keyRepo, repo)
}
