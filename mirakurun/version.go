// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"io"
	"net/http"
)

// Version represents a Mirakurun version.
type Version struct {
	Current string `json:"current"`
	Latest  string `json:"latest"`
}

// CheckVersion fetches a Mirakurun version.
func (c *Client) CheckVersion(ctx context.Context) (*Version, *http.Response, error) {
	req, err := c.NewRequest("GET", "version", nil)
	if err != nil {
		return nil, nil, err
	}

	version := new(Version)
	resp, err := c.Do(ctx, req, version)
	if err != nil {
		return nil, resp, err
	}

	return version, resp, nil
}

// VersionUpdateOptions specifies the optional parameters to the Client.UpdateVersion method.
type VersionUpdateOptions struct {
	Force bool `url:"force,omitempty"`
}

// UpdateVersion updates a Mirakurun version.
func (c *Client) UpdateVersion(ctx context.Context, opt *VersionUpdateOptions) (io.ReadCloser, *http.Response, error) {
	u, err := addOptions("version/update", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("PUT", u, nil)
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
