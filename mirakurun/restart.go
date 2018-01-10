// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"net/http"
)

// RestartService ...
type RestartService service

// RestartResponse ...
type RestartResponse struct {
	PID int `json:"_cmd_pid"`
}

// Do ...
func (s *RestartService) Do(ctx context.Context) (*RestartResponse, *http.Response, error) {
	req, err := s.client.NewRequest("PUT", "restart", nil)
	if err != nil {
		return nil, nil, err
	}

	restart := new(RestartResponse)
	resp, err := s.client.Do(ctx, req, restart)
	if err != nil {
		return nil, resp, err
	}

	return restart, resp, nil
}
