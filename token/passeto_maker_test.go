package token

import (
	"simple_bank_app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPassetoMaker(t *testing.T) {
	maker, err := NewPassetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.WithinDuration(t, expiredAt, payload.ExpiresAtAt, time.Second)
}

func TestExpiredPassetoToken(t *testing.T) {
	maker, err := NewPassetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

// func TestInvalidJWTTOKENlgNone(t *testing.T) {
// 	payload, err := NewPayload(util.RandomOwner(), time.Minute)
// 	require.NoError(t, err)

// 	jwtTokem := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
// 	token, err := jwtTokem.SignedString(jwt.UnsafeAllowNoneSignatureType)
// 	require.NoError(t, err)

// 	maker, err := NewJWTMaker(util.RandomString(32))
// 	require.NoError(t, err)

// 	payload, err = maker.VerifyToken(token)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrInvalidToken.Error())
// 	require.Nil(t, payload)
// }
