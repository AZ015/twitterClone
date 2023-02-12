//go:build integration
// +build integration

package domain

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"twitter"
	"twitter/faker"
	"twitter/test_helpers"
)

func TestIntegrationAuthService_Register(t *testing.T) {
	validInput := twitter.RegisterInput{
		Username:        faker.Username(),
		Email:           faker.Email(),
		Password:        faker.Password,
		ConfirmPassword: faker.Password,
	}
	t.Run("can register a user", func(t *testing.T) {
		ctx := context.Background()

		defer test_helpers.TeardownDB(ctx, t, db)

		res, err := authService.Register(ctx, validInput)
		require.NoError(t, err)

		require.NotEmpty(t, res.AccessToken)
		require.NotEmpty(t, res.User.ID)
		require.NotEmpty(t, res.User.Email)
		require.NotEmpty(t, res.User.Username)
	})

	t.Run("existing username", func(t *testing.T) {
		ctx := context.Background()

		defer test_helpers.TeardownDB(ctx, t, db)

		_, err := authService.Register(ctx, validInput)
		require.NoError(t, err)

		_, err = authService.Register(ctx, twitter.RegisterInput{
			Username:        validInput.Username,
			Email:           faker.Email(),
			Password:        faker.Password,
			ConfirmPassword: faker.Password,
		})

		require.ErrorIs(t, err, twitter.ErrUsernameTaken)
	})
}
