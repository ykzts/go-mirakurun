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

// LogService ...
type LogService service

// Get ...
func (s *LogService) Get(ctx context.Context) (*bytes.Buffer, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "log", nil)
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(ctx, req, buf)
	if err != nil {
		return nil, resp, err
	}

	return buf, resp, nil
}

// GetStream ...
func (s *LogService) GetStream(ctx context.Context, w io.Writer) (*http.Response, error) {
	req, err := s.client.NewRequest("GET", "log/stream", nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
