package awssecrets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSecretManager(t *testing.T) {
	t.SkipNow()

	got, err := NewSecretManager()
	require.NoError(t, err)
	require.NotNil(t, got)
}
