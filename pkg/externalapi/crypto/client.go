package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type AppID string

const (
	defaultHost     = "api.nomics.com"
	defaultBasePath = "v1"
	queryApi        = "key"
)

type Client struct {
	appID    AppID
	service  *url.URL
	basePath string
}

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
	absUrl, err := c.buildUrl(path, queryArgs)
	if err != nil {
		return err
	}
	log.Println("URL", absUrl.String())
	response, err := http.Get(absUrl.String())
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

func (c *Client) buildUrl(path string, queryArgs map[string]string) (*url.URL, error) {
	params := url.Values{}
	params.Add(queryApi, string(c.appID))
	for k, v := range queryArgs {
		params.Add(k, v)
	}
	rawUrl := fmt.Sprintf("https://%s/%s/%s?%s",
		c.service.String(),	c.basePath, path, params.Encode())
	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	absUrl := c.service.ResolveReference(parsed)
	return absUrl, nil
}
