package twitter

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUserIDFromContext(t *testing.T) {
	t.Run("get user id from context", func(t *testing.T) {
		ctx := context.Background()

		ctx = context.WithValue(ctx, contextAuthIDKey, "123")

		userId, err := GetUserIdFromContext(ctx)
		require.NoError(t, err)
		require.Equal(t, "123", userId)
	})

	t.Run("return err if not user id", func(t *testing.T) {
		ctx := context.Background()

		_, err := GetUserIdFromContext(ctx)
		require.ErrorIs(t, err, ErrNoUserIDInContext)
	})

	t.Run("return err if value is not a string", func(t *testing.T) {
		ctx := context.Background()

		ctx = context.WithValue(ctx, contextAuthIDKey, 123)

		_, err := GetUserIdFromContext(ctx)
		require.ErrorIs(t, err, ErrNoUserIDInContext)
	})
}

func TestPutUserIDIntoContext(t *testing.T) {
	t.Run("add user id into context", func(t *testing.T) {
		ctx := context.Background()

		ctx = PutUserIDIntoContext(ctx, "123")

		userId, err := GetUserIdFromContext(ctx)
		require.NoError(t, err)
		require.Equal(t, "123", userId)
	})
}
