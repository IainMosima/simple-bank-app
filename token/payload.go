package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Returned errors
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains yhe payload of the token
type Payload struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	IssuedAt    time.Time `json:"issued_at"`
	ExpiresAtAt time.Time `json:"expires_at"`
}

// NewPayload a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:          tokenId,
		Username:    username,
		IssuedAt:    time.Now(),
		ExpiresAtAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAtAt) {
		return ErrExpiredToken
	}
	return nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiresAtAt), nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

func (payload *Payload) GetIssuer() (string, error) {
	return "", nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (payload *Payload) GetSubject() (string, error) {
	return "", nil
}
