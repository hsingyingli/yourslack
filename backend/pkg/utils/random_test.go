package utils

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	var number int64
	var err error

	for i := 0; i < 100000; i++ {
		min := rand.Int63()
		max := rand.Int63()
		number, err = RandomInt(min, max)

		if min > max {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.GreaterOrEqual(t, number, min)
			require.LessOrEqual(t, number, max)
		}
	}
}

func TestRandomString(t *testing.T) {
	var s string
	var err error

	for i := 0; i < 1000; i++ {
		length := rand.Int63n(32)
		s, err = RandomString(int(length))

		require.NoError(t, err)
		require.Len(t, s, int(length))
	}
}
