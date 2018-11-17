package wrike

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://www.wrike.com/"
	apiVersionPath = "api/v4/"
	userAgent      = "go-wrike"
)

// Client Wrike API client
type Client struct {
	client    *http.Client
	baseURL   *url.URL
	token     string
	userAgent string
	Folders   *FolderService
	Timelogs  *TimelogService
}

// NewClient return client
func NewClient(httpClient *http.Client, token string) *Client {
	client := newClient(httpClient)
	client.token = token
	return client
}

// SetBaseURL sets the base URL for API requests to a custom endpoint. urlStr
// should always be specified with a trailing slash.
func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

// NewRequest build request
func (c *Client) NewRequest(method string, path string) (*http.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	req := &http.Request{
		Method:     method,
		URL:        &u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	return req, nil
}

// Response struct
type Response struct {
	*http.Response
}

// Do execute a request
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

// CheckResponse check response returned
func CheckResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return nil
}

func newResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}
	return response
}

func newClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{client: httpClient, userAgent: userAgent}
	if err := c.SetBaseURL(defaultBaseURL); err != nil {
		panic(err)
	}

	c.Folders = &FolderService{client: c}
	c.Timelogs = &TimelogService{client: c}

	return c
}
