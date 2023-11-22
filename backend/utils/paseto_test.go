package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("12345678901234567890123456789012")
	require.NoError(t, err)

	token, err := maker.Create("Bob", time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, ok := maker.Verify(token)
	require.True(t, ok)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, "Bob", payload.Username)
	require.WithinDuration(t, time.Now(), payload.IssuedAt, time.Second)
	require.WithinDuration(t, time.Now().Add(time.Minute), payload.ExpiredAt, time.Second)

	// Check for expired
	token, err = maker.Create("Bob", time.Millisecond*100)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, ok = maker.Verify(token)
	require.True(t, ok)
	require.NotEmpty(t, payload)
	time.Sleep(time.Millisecond * 200)
	payload, ok = maker.Verify(token)
	require.False(t, ok)
	require.Empty(t, payload)

	payload, ok = maker.Verify(nil)
	require.False(t, ok)
	require.Empty(t, payload)

	// Check other errors
	token, err = maker.Create("", time.Second)
	require.Error(t, err)
	require.Empty(t, token)

	token, err = maker.Create("Bob", -time.Second)
	require.Error(t, err)
	require.Empty(t, token)

	maker, err = NewPasetoMaker("1234567890123456789012345678901")
	require.Error(t, err)
	require.Empty(t, maker)
}
