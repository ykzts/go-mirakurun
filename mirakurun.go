/*
 * Copyright (c) 2018 Yamagishi Kazutoshi
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package mirakurun // import "ykzts.com/x/mirakurun"

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
	defaultBaseURL   = "http://127.0.0.1:40772/api/"
	defaultUserAgent = "go-mirakurun/1.0"
)

// A Client manages communication with the Mirakurun API.
type Client struct {
	client *http.Client

	BaseURL   *url.URL
	Priority  int
	UserAgent string
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

	return &Client{client: httpClient, BaseURL: baseURL}
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("mirakurun: BaseURL must have a trailing slash, but %q does not", c.BaseURL)
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

	if c.UserAgent == "" {
		req.Header.Set("User-Agent", defaultUserAgent)
	} else {
		req.Header.Set("User-Agent", fmt.Sprintf("%s %s", c.UserAgent, defaultBaseURL))
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

	if resp.StatusCode < 200 || resp.StatusCode > 202 {
		return resp, fmt.Errorf("mirakurun: %s", resp.Status)
	}

	if v != nil {
		dec := json.NewDecoder(resp.Body)
		dec.Decode(v)
	}

	return resp, nil
}

func (c *Client) requestStream(ctx context.Context, method string, u string) (io.ReadCloser, *http.Response, error) {
	req, err := c.NewRequest(method, u, nil)
	if err != nil {
		return nil, nil, err
	}

	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, resp, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 202 {
		resp.Body.Close()
		return nil, resp, fmt.Errorf("mirakurun: %s", resp.Status)
	}

	return resp.Body, resp, nil
}

// DecodeOptions specifies the optional parameters to various Stream method that support decoding.
type DecodeOptions struct {
	Decode int `url:"decode,omitempty"`
}

func (c *Client) getTS(ctx context.Context, u string, decode bool) (io.ReadCloser, *http.Response, error) {
	opt := &DecodeOptions{Decode: 0}
	if decode {
		opt.Decode = 1
	}

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	stream, resp, err := c.requestStream(ctx, "GET", u)
	if err != nil {
		return nil, resp, err
	}

	if contentType := resp.Header.Get("Content-Type"); contentType != "video/MP2T" {
		stream.Close()
		return nil, resp, fmt.Errorf("mirakurun: invalid content type, but %s does not", contentType)
	}

	return stream, resp, nil
}
