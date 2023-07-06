package controller

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Write unit tests for function NewServer in server.go. You should test that the server is created correctly and that it has the correct address and handler. You should also test that the correct error is returned if the server cannot be created
func TestNewServer(t *testing.T) {
	tests := map[string]struct {
		want        *http.Server
		expectedErr error
	}{
		"success": {
			want: &http.Server{
				Addr:    ":8080",
				Handler: http.NewServeMux(),
			},
			expectedErr: nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewServer()
			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.want.Addr, got.Addr)
		})
	}
}
