package token

import (
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	maker := NewJWTMaker("secret")
	duration := time.Minute

	// Adjusted assignment to match the return values
	token, err := maker.CreateToken("user_id", duration)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	if token == "" {
		t.Fatal("expected a token to be created")
	}
}