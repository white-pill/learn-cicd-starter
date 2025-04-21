package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		want    string
		wantErr bool
	}{
		"no authorization header":                  {headers: http.Header{}, want: "", wantErr: true},
		"authorization header no white space":      {headers: http.Header{"Authorization": {"ApiKey123"}}, want: "", wantErr: true},
		"authorization header with correct ApiKey": {headers: http.Header{"Authorization": {"ApiKey 123"}}, want: "123", wantErr: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got none (value: %v)", got)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got != tc.want {
					t.Fatalf("expected: %v, got: %v", tc.want, got)
				}
			}
		})
	}
}
