package twitter

import "errors"

var (
	ErrValidation         = errors.New("validation error")
	ErrNotFound           = errors.New("not found error")
	ErrBadCredentials     = errors.New("err/password wrong combination")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrNoUserIDInContext  = errors.New("no user id in context")
	ErrGenAccessToken     = errors.New("generate access token error")
	ErrUnauthenticated    = errors.New("unauthenticated error")
	ErrInvalidUUID        = errors.New("invalid uuid")
	ErrForbidden          = errors.New("forbidden")
)
