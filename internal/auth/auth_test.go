package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

// TestGetAPIKey - tests the GetAPIKey function
func TestGetApiKey(t *testing.T) {
	type test struct {
		name        string
		headers     http.Header
		expected    string
		expectedErr error
	}

	tests := []test{
		{
			name:        "Valid API Key",
			headers:     http.Header{"Authorization": []string{"ApiKey abc123"}},
			expected:    "abc123",
			expectedErr: nil},
		{name: "No Auth Header",
			headers:     http.Header{},
			expected:    "",
			expectedErr: ErrNoAuthHeaderIncluded},
		{name: "Malformed Auth Header",
			headers:     http.Header{"Authorization": []string{"Bearer abc123"}},
			expected:    "",
			expectedErr: errors.New("malformed authorization header")},
		{name: "Empty Auth Header",
			headers:     http.Header{"Authorization": []string{""}},
			expected:    "",
			expectedErr: ErrNoAuthHeaderIncluded},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.headers)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Fatalf("expected: %v, got: %v, Error: %v", tc.expected, got, err)
		}
	}
}
