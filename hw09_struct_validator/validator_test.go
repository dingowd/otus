package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:12"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:abc,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "Kukaracha_id",
				Name:   "Jango",
				Age:    36,
				Email:  "example@mail.ru",
				Role:   "admin",
				Phones: []string{"79519878541"},
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "seven",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 404,
				Body: "Simple",
			},
			expectedErr: ValidatingStructError,
		},
		{
			in: Token{
				Header:    []byte("no response"),
				Payload:   []byte("empty"),
				Signature: []byte("empty"),
			},
			expectedErr: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			require.Equal(t, tt.expectedErr, Validate(tt.in))
			_ = tt
		})
	}
}
