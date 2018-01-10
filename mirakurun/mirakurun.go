package mirakurun

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL   = "http://localhost:40772/api/"
	defaultUserAgent = "go-mirakurun/1"
)

// A Client manages communication with the Mirakurun API.
type Client struct {
	client *http.Client

	BaseURL   *url.URL
	Priority  int
	UserAgent string

	common service

	Channels *ChannelsService
	Services *ServicesService
	Tuners   *TunersService
	Events   *EventsService
	Config   *ConfigService
	Log      *LogService
	Version  *VersionService
	Status   *StatusService
	Restart  *RestartService
}

type service struct {
	client *Client
}

func addOptions(s string, opt interface{}) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

// NewClient returns a new Mirakurun API client.
func NewClient() *Client {
	httpClient := &http.Client{}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c
	c.Channels = (*ChannelsService)(&c.common)
	c.Services = (*ServicesService)(&c.common)
	c.Tuners = (*TunersService)(&c.common)
	c.Events = (*EventsService)(&c.common)
	c.Config = (*ConfigService)(&c.common)
	c.Log = (*LogService)(&c.common)
	c.Version = (*VersionService)(&c.common)
	c.Status = (*StatusService)(&c.common)
	c.Restart = (*RestartService)(&c.common)

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	req.Header.Set("X-Mirakurun-Priority", strconv.Itoa(c.Priority))

	if body != nil {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	return req, nil
}

// Do sends an API requests and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			dec := json.NewDecoder(resp.Body)
			dec.Decode(v)
		}
	}

	return resp, nil
}
