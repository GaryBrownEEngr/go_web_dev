package user

import (
	"testing"

	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
	"github.com/stretchr/testify/require"
)

func Test_hashPassword(t *testing.T) {
	hash1, err := hashPassword("bacon1234")
	require.NoError(t, err)
	hash2, err := hashPassword("bacon1234")
	require.NoError(t, err)
	hash3, err := hashPassword("bacon1234")
	require.NoError(t, err)

	require.NotEqual(t, hash1, hash2)
	require.NotEqual(t, hash1, hash3)
	require.NotEqual(t, hash2, hash3)

	require.True(t, doPasswordsMatch(hash1, "bacon1234"))
	require.True(t, doPasswordsMatch(hash2, "bacon1234"))
	require.True(t, doPasswordsMatch(hash3, "bacon1234"))

	require.False(t, doPasswordsMatch(hash1, "bacon12345"))
	require.False(t, doPasswordsMatch(hash2, "bacon123"))
	require.False(t, doPasswordsMatch(hash3, "bacon1235"))
	require.False(t, doPasswordsMatch(hash3, ""))

	require.False(t, doPasswordsMatch("", "bacon1235"))
	require.False(t, doPasswordsMatch("", ""))
	require.False(t, doPasswordsMatch("abc", ""))
}

func TestNewUserStore(t *testing.T) {
	t.SkipNow()

	s, err := NewUserStore()
	require.NoError(t, err)

	user, err := s.VerifyPassword("Bacon", "12345678")
	require.Error(t, err)
	require.Nil(t, user)

	user, err = s.CreateUser("", "")
	require.Error(t, err)
	require.Nil(t, user)
	user, err = s.CreateUser("abc_test", "")
	require.Error(t, err)
	require.Nil(t, user)

	user, err = s.CreateUser("abc_test", "testPassword")
	require.NoError(t, err, utils.UnwrapAllErrorsForLog(err))
	require.NotNil(t, user)

	user2, err := s.CreateUser("abc_test", "testPassword")
	require.Error(t, err)
	require.Nil(t, user2)

	user, err = s.VerifyPassword("abc_test", "testPassword")
	require.NoError(t, err)
	require.NotNil(t, user)
	user, err = s.VerifyPassword("abc_test", "testPassword1")
	require.Error(t, err)
	require.Nil(t, user)
}
