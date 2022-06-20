package crypto

import (
	"fmt"
	"github.com/nwillc/cryptoport/gjson"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/result"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// AppID API key
type AppID string

const (
	defaultAppIDEnv = "NOMICS_APP_ID"
	defaultHost     = "api.nomics.com"
	defaultBasePath = "v1"
	queryAPI        = "key"
)

// Client used to access services.
type Client struct {
	appID    AppID
	service  *url.URL
	basePath string
}

// DefaultAppID returns AppID based on NOMICS_APP_ID env variable.
func DefaultAppID() AppID {
	return AppID(os.Getenv(defaultAppIDEnv))
}

// NewClient creates a new Client using given AppID.
func NewClient(appID AppID) *genfuncs.Result[*Client] {
	service := genfuncs.NewResultError(url.Parse(defaultHost))
	return result.Map(service, func(url *url.URL) *genfuncs.Result[*Client] {
		return genfuncs.NewResult(&Client{
			appID:    appID,
			service:  url,
			basePath: defaultBasePath,
		})
	})
}

func (c *Client) getTickerInfo(path string, queryArgs container.GMap[string, string]) *genfuncs.Result[*container.GSlice[TickerInfo]] {
	absURL := c.buildURL(path, queryArgs)
	response := result.Map(absURL, func(url *url.URL) *genfuncs.Result[*http.Response] {
		return genfuncs.NewResultError(http.Get(url.String()))
	})
	body := result.Map(response, func(response *http.Response) *genfuncs.Result[[]byte] {
		if response.StatusCode != 200 {
			return genfuncs.NewError[[]byte](fmt.Errorf("http response %d", response.StatusCode))
		}
		return genfuncs.NewResultError(ioutil.ReadAll(response.Body))
	})
	return result.Map(body, func(bytes []byte) *genfuncs.Result[*container.GSlice[TickerInfo]] {
		return gjson.Unmarshal[container.GSlice[TickerInfo]](bytes)
	})
}

func (c *Client) buildURL(path string, queryArgs container.GMap[string, string]) *genfuncs.Result[*url.URL] {
	params := url.Values{}
	params.Add(queryAPI, string(c.appID))
	for k, v := range queryArgs {
		params.Add(k, v)
	}
	rawURL := fmt.Sprintf("https://%s/%s/%s?%s",
		c.service.String(), c.basePath, path, params.Encode())
	parsed := genfuncs.NewResultError(url.Parse(rawURL))
	return result.Map(parsed, func(url *url.URL) *genfuncs.Result[*url.URL] {
		return genfuncs.NewResult(c.service.ResolveReference(url))
	})
}
