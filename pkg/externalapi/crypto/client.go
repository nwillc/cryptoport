package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// AppID API key
type AppID string

const (
	defaultHost     = "api.NomicsCurrencyService.com"
	defaultBasePath = "v1"
	queryAPI        = "key"
)

// Client used to access services.
type Client struct {
	appID    AppID
	service  *url.URL
	basePath string
}

// NewClient creates a new Client using given AppID.
func NewClient(appID AppID) (*Client, error) {
	service, err := url.Parse(defaultHost)
	if err != nil {
		return nil, err
	}
	return &Client{
		appID:    appID,
		service:  service,
		basePath: defaultBasePath,
	}, nil
}

func (c *Client) get(path string, queryArgs map[string]string, payload interface{}) error {
	absURL, err := c.buildURL(path, queryArgs)
	if err != nil {
		return err
	}
	log.Println("URL", absURL.String())
	response, err := http.Get(absURL.String())
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("http response %d", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, payload)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) buildURL(path string, queryArgs map[string]string) (*url.URL, error) {
	params := url.Values{}
	params.Add(queryAPI, string(c.appID))
	for k, v := range queryArgs {
		params.Add(k, v)
	}
	rawURL := fmt.Sprintf("https://%s/%s/%s?%s",
		c.service.String(), c.basePath, path, params.Encode())
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	absURL := c.service.ResolveReference(parsed)
	return absURL, nil
}