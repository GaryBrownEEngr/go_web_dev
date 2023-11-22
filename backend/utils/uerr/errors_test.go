package uerr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnwrapAllErrorsForLog(t *testing.T) {
	s := UnwrapAllErrorsForLog(nil)
	require.Empty(t, s)

	err := fmt.Errorf("abc")
	s = UnwrapAllErrorsForLog(err)
	require.Equal(t, "abc", s)

	err = fmt.Errorf("123-%w", err)
	s = UnwrapAllErrorsForLog(err)
	require.Equal(t, "123-abc :: abc", s)
}

func TestHashError(t *testing.T) {
	err1 := fmt.Errorf("abc")
	err2 := fmt.Errorf("abb")
	h1 := HashError(err1)
	h2 := HashError(err2)

	require.Equal(t, "BA7816BF8F01", h1)
	require.Equal(t, "715EDF8BA872", h2)
}
