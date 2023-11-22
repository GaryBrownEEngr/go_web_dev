package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hashPassword(t *testing.T) {
	hash1, err := HashPassword("bacon1234")
	require.NoError(t, err)
	hash2, err := HashPassword("bacon1234")
	require.NoError(t, err)
	hash3, err := HashPassword("bacon1234")
	require.NoError(t, err)

	require.NotEqual(t, hash1, hash2)
	require.NotEqual(t, hash1, hash3)
	require.NotEqual(t, hash2, hash3)

	require.True(t, VerifyPassword(hash1, "bacon1234"))
	require.True(t, VerifyPassword(hash2, "bacon1234"))
	require.True(t, VerifyPassword(hash3, "bacon1234"))

	require.False(t, VerifyPassword(hash1, "bacon12345"))
	require.False(t, VerifyPassword(hash2, "bacon123"))
	require.False(t, VerifyPassword(hash3, "bacon1235"))
	require.False(t, VerifyPassword(hash3, ""))

	require.False(t, VerifyPassword("", "bacon1235"))
	require.False(t, VerifyPassword("", ""))
	require.False(t, VerifyPassword("abc", ""))
}
