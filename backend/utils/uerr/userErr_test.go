package uerr

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUErr(t *testing.T) {
	err := UErr("abc", 200)
	require.Equal(t, "abc", err.Error())
	require.Equal(t, nil, errors.Unwrap(err))

	var err2 *UserErrorData
	ok := errors.As(err, &err2)
	require.True(t, ok)
	msg, code := err2.UserMsgAndCode()
	require.Equal(t, "abc", msg)
	require.Equal(t, 200, code)
	require.Equal(t, 200, err2.UserCode())
	require.Equal(t, false, err2.ShouldLog())
}

func TestUErrLogHash(t *testing.T) {
	err := UErrLogHash("abc", 400, fmt.Errorf("inner"))
	require.Equal(t, "abc 33BF6FBD7CD8", err.Error())
	innerError := errors.Unwrap(err)
	require.NotNil(t, innerError)
	require.Equal(t, "inner", innerError.Error())

	var err2 *UserErrorData
	ok := errors.As(err, &err2)
	require.True(t, ok)
	msg, code := err2.UserMsgAndCode()
	require.Equal(t, "abc 33BF6FBD7CD8", msg)
	require.Equal(t, 400, code)
	require.Equal(t, 400, err2.UserCode())
	require.Equal(t, true, err2.ShouldLog())
}
