package domain

import (
	"context"
	"twitter"
	"twitter/uuid"
)

type TweetService struct {
	TweetRepo twitter.TweetRepo
}

func NewTweetService(tr twitter.TweetRepo) *TweetService {
	return &TweetService{
		TweetRepo: tr,
	}
}

func (t *TweetService) All(ctx context.Context) ([]twitter.Tweet, error) {
	return t.TweetRepo.All(ctx)
}

func (t *TweetService) Create(ctx context.Context, in twitter.CreateTweetInput) (twitter.Tweet, error) {
	currentUserID, err := twitter.GetUserIdFromContext(ctx)
	if err != nil {
		return twitter.Tweet{}, twitter.ErrUnauthenticated
	}

	in.Sanitize()

	if err := in.Validate(); err != nil {
		return twitter.Tweet{}, err
	}

	tweet, err := t.TweetRepo.Create(ctx, twitter.Tweet{
		Body:   in.Body,
		UserID: currentUserID,
	})
	if err != nil {
		return twitter.Tweet{}, err
	}

	return tweet, nil
}

func (t *TweetService) GetByID(ctx context.Context, id string) (twitter.Tweet, error) {
	if !uuid.Validate(id) {
		return twitter.Tweet{}, twitter.ErrInvalidUUID
	}
	return t.TweetRepo.GetByID(ctx, id)
}

func (t TweetService) Delete(ctx context.Context, id string) error {
	currentUserID, err := twitter.GetUserIdFromContext(ctx)
	if err != nil {
		return twitter.ErrUnauthenticated
	}

	if !uuid.Validate(id) {
		return twitter.ErrInvalidUUID
	}

	tweet, err := t.TweetRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if !tweet.CanDelete(twitter.User{ID: currentUserID}) {
		return twitter.ErrForbidden
	}

	return t.TweetRepo.Delete(ctx, id)
}
