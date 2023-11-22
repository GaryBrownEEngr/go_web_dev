package stacktrs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	err := Wrap(fmt.Errorf("abc"))
	require.NotNil(t, err)
	msg := err.Error()
	require.Contains(t, msg, "go_web_dev/backend/utils/stacktrs.wrap")

	errInner := errors.Unwrap(err)
	require.NotNil(t, errInner)
	require.Equal(t, "abc", errInner.Error())

	err = Wrap(nil)
	require.Nil(t, err)
}

func TestErrorf(t *testing.T) {
	err := Errorf("abc %d", 12)
	require.NotNil(t, err)
	msg := err.Error()
	require.Contains(t, msg, "abc 12")
	require.Contains(t, msg, "go_web_dev/backend/utils/stacktrs.wrap")

	err2 := Wrap(err)
	require.NotNil(t, err)
	require.Equal(t, err, err2)
}
