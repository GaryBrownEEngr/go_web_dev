package sessionuser

import (
	"testing"

	"github.com/GaryBrownEEngr/go_web_dev/backend/utils/uerr"
	"github.com/stretchr/testify/require"
)

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
	require.NoError(t, err, uerr.UnwrapAllErrorsForLog(err))
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
