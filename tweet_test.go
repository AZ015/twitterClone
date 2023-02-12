package twitter

import (
	"github.com/stretchr/testify/require"
	"testing"
	"twitter/faker"
)

func TestCreateTweetInput_Sanitize(t *testing.T) {
	in := CreateTweetInput{
		Body: " hello      ",
	}

	want := CreateTweetInput{
		Body: "hello",
	}

	in.Sanitize()

	require.Equal(t, in, want)
}

func TestCreateTweetInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input CreateTweetInput
		err   error
	}{
		{
			name: "valid",
			input: CreateTweetInput{
				Body: "hello",
			},
			err: nil,
		},
		{
			name: "tweet not long enough",
			input: CreateTweetInput{
				Body: "h",
			},
			err: nil,
		},
		{
			name: "tweet too long",
			input: CreateTweetInput{
				Body: faker.RandStr(1000),
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
