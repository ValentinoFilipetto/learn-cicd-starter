package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	ExpectedAPIKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9c"
	headers := http.Header{
		"Authorization": []string{"ApiKey eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9c"},
	}

	APIKey, _ := GetAPIKey(headers)

	if ExpectedAPIKey != APIKey {
		t.Fatalf("Expected key %s, got %s", ExpectedAPIKey, APIKey)
	}
}
func TestGetAPIKey_ValidHeader(t *testing.T) {
	expected := "testkey123"
	headers := http.Header{
		"Authorization": []string{"ApiKey " + expected},
	}
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if key != expected {
		t.Fatalf("Expected key %s, got %s", expected, key)
	}
}
