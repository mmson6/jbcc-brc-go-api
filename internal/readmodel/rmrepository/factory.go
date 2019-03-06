package rmrepository

import (
	"context"

	"code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/readmodel/rmdynamodb"
)

type ctxKey string

const keyRepo ctxKey = "rm-repo"

func Build(ctx context.Context) Repository {
	val := ctx.Value(keyRepo)
	if tbl, ok := val.(Repository); ok {
		return tbl
	} else {
		tbl := rmdynamodb.BuildTable(ctx)
		return New(tbl)
	}
}

func Inject(ctx context.Context, repo Repository) context.Context {
	return context.WithValue(ctx, keyRepo, repo)
}
