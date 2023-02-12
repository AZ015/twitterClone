//go:build integration
// +build integration

package domain

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"testing"
	"twitter"
	"twitter/config"
	"twitter/jwt"
	"twitter/postgres"
)

var (
	conf             *config.Config
	db               *postgres.DB
	authService      twitter.AuthService
	authTokenService twitter.AuthTokenService
	tweetService     twitter.TweetService
	userRepo         twitter.UserRepo
	tweetRepo        twitter.TweetRepo
)

func TestMain(t *testing.M) {
	ctx := context.Background()

	config.LoadEnv(".env.test")

	passwordCost = bcrypt.MinCost

	conf = config.New()

	db = postgres.New(ctx, conf)
	defer db.Close()

	if err := db.Drop(); err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	userRepo = postgres.NewUserRepo(db)
	tweetRepo = postgres.NewTweetRepo(db)
	authTokenService = jwt.NewTokenService(conf)

	authService = NewAuthService(userRepo, authTokenService)
	tweetService = NewTweetService(tweetRepo)

	os.Exit(t.Run())
}
