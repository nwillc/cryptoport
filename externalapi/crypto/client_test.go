package crypto

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("foo")
	require.True(t, client.Ok())
	assert.NotEmpty(t, client.OrEmpty().appID)
	assert.NotEmpty(t, client.OrEmpty().service.String())
	assert.NotEmpty(t, client.OrEmpty().basePath)
}

func TestClient_buildUrl(t *testing.T) {
	client := NewClient("foo")
	require.True(t, client.Ok())
	url := client.OrEmpty().buildURL("test", map[string]string{"now": "later"})
	require.True(t, url.Ok())
	assert.Equal(t, "https://api.nomics.com/v1/test?key=foo&now=later", url.String())
}
