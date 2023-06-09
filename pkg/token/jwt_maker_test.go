package token

import (
	"testing"
	"time"

	"github.com/Damione1/portfolio-playground/util"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	email := util.RandomEmail()
	duration := time.Minute

	IssuedTime := time.Now()
	ExpireTime := time.Now().Add(duration)

	token, _, err := maker.CreateToken(email, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, email, payload.Email)
	require.WithinDuration(t, IssuedTime, payload.IssuedTime, time.Second)
	require.WithinDuration(t, ExpireTime, payload.ExpireTime, time.Second)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {

	payload, err := NewPayload(util.RandomEmail(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.ValidateToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)

}
