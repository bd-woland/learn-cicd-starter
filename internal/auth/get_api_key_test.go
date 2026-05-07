package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input  http.Header
		result string
		err    string
	}{
		"simple1": {input: http.Header{
			"Authorization": []string{"ApiKey api_key"},
		}, result: "api_key", err: ""},
		"simple2": {input: http.Header{
			"Authorization": []string{"ApiKey aPi_kEy"},
		}, result: "aPi_kEy", err: ""},
		"multipe values": {input: http.Header{
			"Authorization": []string{"ApiKey first", "ApiKey second"},
		}, result: "first", err: ""},
		"malformed": {input: http.Header{
			"Authorization": []string{"Bearer api_key"},
		}, result: "", err: "malformed authorization header"},
		"missing": {input: http.Header{
			"Accept": []string{"application/json"},
		}, result: "", err: "no authorization header included"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.input)

			if err != nil {
				if tc.result != "" {
					t.Fatalf("expected: %v, got error: %v", tc.result, err)
				} else if err.Error() != tc.err {
					t.Fatalf("expected error message to be: %v, got: %v", tc.err, err.Error())
				}
			} else if tc.err != "" {
				t.Fatalf("expected error: %v, got result: %v", tc.err, key)
			} else if tc.result != key {
				t.Fatalf("expected: %v, got: %v", tc.result, key)
			}
		})
	}
}
