package crypto

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("foo")
	require.NoError(t, err)
	require.NotNil(t, client)
	assert.NotEmpty(t, client.appID)
	assert.NotEmpty(t, client.service.String())
	assert.NotEmpty(t, client.basePath)
}

func TestClient_buildUrl(t *testing.T) {
	client, err := NewClient("foo")
	require.NoError(t, err)
	url, err := client.buildUrl("test", map[string]string{"now": "later"})
	require.NoError(t, err)
	assert.Equal(t, "https://api.nomics.com/v1/test?key=foo&now=later", url)
}
