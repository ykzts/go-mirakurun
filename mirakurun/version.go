// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"io"
	"net/http"
)

// VersionService ...
type VersionService service

// Version ...
type Version struct {
	Current string `json:"current"`
	Latest  string `json:"latest"`
}

// Get ...
func (s *VersionService) Get(ctx context.Context) (*Version, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "version", nil)
	if err != nil {
		return nil, nil, err
	}

	version := new(Version)
	resp, err := s.client.Do(ctx, req, version)
	if err != nil {
		return nil, resp, err
	}

	return version, resp, nil
}

// VersionUpdateOptions ...
type VersionUpdateOptions struct {
	Force bool `url:"force,omitempty"`
}

// Update ...
func (s *VersionService) Update(ctx context.Context, opt *VersionUpdateOptions, w io.Writer) (*http.Response, error) {
	u, err := addOptions("version/update", opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
