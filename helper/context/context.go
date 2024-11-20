package context

import (
	"context"
)

func CreateContext() context.Context {
	return context.Background()
}

func SetTokenStructToContext(ctx context.Context, key string, token interface{}) context.Context {
	return context.WithValue(ctx, key, token)
}
