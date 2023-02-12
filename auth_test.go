package twitter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterInput_Validate(t *testing.T) {
	tests := []struct {
		name string
		in   RegisterInput
		err  error
	}{
		{
			name: "valid",
			in: RegisterInput{
				Username:        "bob",
				Email:           "bob@email.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		{
			name: "too short username",
			in: RegisterInput{
				Username:        "b",
				Email:           "bob@",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "invalid email",
			in: RegisterInput{
				Username:        "bob",
				Email:           "bob@",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "too short password",
			in: RegisterInput{
				Username:        "bob",
				Email:           "bob@mail.com",
				Password:        "passw",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "password not equal confirm password",
			in: RegisterInput{
				Username:        "bob",
				Email:           "bob@mail.com",
				Password:        "password!",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()

			if tt.err != err {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRegisterInput_Sanitize(t *testing.T) {
	in := RegisterInput{
		Username:        " bob ",
		Email:           " BOB@gmail.com  ",
		Password:        "password",
		ConfirmPassword: "password",
	}

	want := RegisterInput{
		Username:        "bob",
		Email:           "bob@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}
	in.Sanitize()

	require.Equal(t, want, in)
}

func TestLoginInput_Sanitize(t *testing.T) {
	in := LoginInput{
		Email:    " BOB@gmail.com  ",
		Password: "password",
	}

	want := LoginInput{
		Email:    "bob@gmail.com",
		Password: "password",
	}
	in.Sanitize()

	require.Equal(t, want, in)
}

func TestLoginInput_Validate(t *testing.T) {
	tests := []struct {
		name string
		in   LoginInput
		err  error
	}{
		{
			name: "valid",
			in: LoginInput{
				Email:    "bob@email.com",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			in: LoginInput{
				Email:    "bob@",
				Password: "password",
			},
			err: ErrValidation,
		},
		{
			name: "empty password",
			in: LoginInput{
				Email:    "bob@mail.com",
				Password: "",
			},
			err: ErrValidation,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()

			if tt.err != err {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
