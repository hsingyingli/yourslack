package token

import (
	"testing"
	"time"

	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
	"github.com/hsingyingli/yourslack-backend/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestNewPasetoMaker(t *testing.T) {
	wrongSymmetricKey, _ := utils.RandomString(31)

	maker, err := NewPasetoMaker(wrongSymmetricKey)
	require.Errorf(t, err, "length of symmetricKey must be greater than 32")
	require.Empty(t, maker)

	symmetricKey, _ := utils.RandomString(32)
	maker, err = NewPasetoMaker(symmetricKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)
}

func TestCreateAndVerifyToken(t *testing.T) {
	symmetricKey, _ := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	userID, _ := utils.RandomInt(1, 100)
	username, _ := utils.RandomString(8)
	email, _ := utils.RandomString(8)

	user := db.User{
		ID:       userID,
		Username: username,
		Email:    email,
	}

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, payload, err := maker.CreateToken(user, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	testPayload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, user.Username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
	require.Equal(t, testPayload.Username, payload.Username)
	require.WithinDuration(t, testPayload.IssuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, testPayload.ExpiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	symmetricKey, _ := utils.RandomString(32)
	maker, err := NewPasetoMaker(symmetricKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	userID, _ := utils.RandomInt(1, 100)
	username, _ := utils.RandomString(8)
	email, _ := utils.RandomString(8)

	user := db.User{
		ID:       userID,
		Username: username,
		Email:    email,
	}
	duration := -time.Minute
	token, payload, err := maker.CreateToken(user, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	testPayload, err := maker.VerifyToken(token)
	require.EqualError(t, err, "token has expired")
	require.Empty(t, testPayload)
}
