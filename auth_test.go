package graphql_twitter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testcase struct {
	name  string
	input RegisterInput
	err   error
}

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []testcase{
		{
			name: "valid",
			input: RegisterInput{
				Email:           "bob@gmail.com",
				Username:        "bob",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			input: RegisterInput{
				Email:           "bob",
				Username:        "bob",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "too short username",
			input: RegisterInput{
				Email:           "bob@gmail.com",
				Username:        "b",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "too short password",
			input: RegisterInput{
				Email:           "bob@gmail.com",
				Username:        "bob",
				Password:        "pass",
				ConfirmPassword: "pass",
			},
			err: ErrValidation,
		},
		{
			name: "confirm password doesn't match password",
			input: RegisterInput{
				Email:           "bob@gmail.com",
				Username:        "bob",
				Password:        "password",
				ConfirmPassword: "wrongpassword",
			},
			err: ErrValidation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				// want to have an error
				require.ErrorIs(t, err, tc.err)
			} else {
				// don't want to have an error
				require.NoError(t, err)
			}

		})
	}
}

func TestRegisterInput_Sanitize(t *testing.T) {
	input := RegisterInput{
		Email:           " BOB@gmail.com ",
		Username:        " bob ",
		Password:        "password",
		ConfirmPassword: "password",
	}
	want := RegisterInput{
		Email:           "bob@gmail.com",
		Username:        "bob",
		Password:        "password",
		ConfirmPassword: "password",
	}

	input.Sanitize()

	require.Equal(t, want, input)
}
