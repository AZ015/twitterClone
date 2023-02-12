package twitter

import (
	"context"
)

var (
	contextAuthIDKey = "currentUserId"
)

func GetUserIdFromContext(ctx context.Context) (string, error) {
	if ctx.Value(contextAuthIDKey) == nil {
		return "", ErrNoUserIDInContext
	}

	userID, ok := ctx.Value(contextAuthIDKey).(string)
	if !ok {
		return "", ErrNoUserIDInContext
	}

	return userID, nil
}

func PutUserIDIntoContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextAuthIDKey, id)
}
