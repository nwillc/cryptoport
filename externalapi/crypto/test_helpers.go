package crypto

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// MightSkipIntegrationTest determines if integraation tests should run.
func MightSkipIntegrationTest(t *testing.T) *Client {
	t.Helper()
	appID := DefaultAppID()
	if appID == "" {
		t.Skipf("integration test env var %s not set", defaultAppIDEnv)
		return nil
	}
	client, err := NewClient(appID)
	require.NoError(t, err)
	return client
}
