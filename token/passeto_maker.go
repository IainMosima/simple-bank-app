package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PassetoMaker struct {
	passeto       *paseto.V2
	symmmetricKey []byte
}

func NewPassetoMaker(symmmetricKey string) (Maker, error) {
	if len(symmmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size")
	}

	maker := &PassetoMaker{
		passeto:       paseto.NewV2(),
		symmmetricKey: []byte(symmmetricKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PassetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return maker.passeto.Encrypt(maker.symmmetricKey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (maker *PassetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.passeto.Decrypt(token, maker.symmmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
