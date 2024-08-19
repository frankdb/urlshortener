package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShortenHandler(t *testing.T) {
	reqBody := []byte(`{"url": "https://www.example.com"}`)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ShortenHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp ShortenResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(resp.ShortURL) == 0 {
		t.Errorf("handler returned empty short URL")
	}
}

func TestGenerateRandomString(t *testing.T) {
	length := 6
	s1 := generateRandomString(length)
	s2 := generateRandomString(length)

	if len(s1) != length {
		t.Errorf("generateRandomString(%d) returned string of length %d, want %d", length, len(s1), length)
	}

	if s1 == s2 {
		t.Errorf("Two calls to generateRandomString(%d) returned identical strings, they should be different", length)
	}

	// Test that only characters from the charset are used
	for _, char := range s1 {
		if !strings.ContainsRune(charset, char) {
			t.Errorf("generateRandomString(%d) used character '%c' which is not in the charset", length, char)
		}
	}
}
