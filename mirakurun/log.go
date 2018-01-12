// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

// GetLog fetches a log.
func (c *Client) GetLog(ctx context.Context) (*bytes.Buffer, *http.Response, error) {
	req, err := c.NewRequest("GET", "log", nil)
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := c.Do(ctx, req, buf)
	if err != nil {
		return nil, resp, err
	}

	return buf, resp, nil
}

// GetLogStream fetches a log stream.
func (c *Client) GetLogStream(ctx context.Context) (io.ReadCloser, *http.Response, error) {
	req, err := c.NewRequest("GET", "log/stream", nil)
	if err != nil {
		return nil, nil, err
	}

	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, resp, err
	}

	return resp.Body, resp, nil
}
