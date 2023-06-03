package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password, _ := RandomString(10)

	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEqual(t, password, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)

	fakePassword, _ := RandomString(10)
	err = CheckPassword(hashedPassword, fakePassword)
	require.Error(t, err)
}
