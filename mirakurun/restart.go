// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"net/http"
)

// RestartResponse represents a Mirakurun restart response.
type RestartResponse struct {
	PID int `json:"_cmd_pid"`
}

// Restart a Mirakurun process.
func (c *Client) Restart(ctx context.Context) (*RestartResponse, *http.Response, error) {
	req, err := c.NewRequest("PUT", "restart", nil)
	if err != nil {
		return nil, nil, err
	}

	restart := new(RestartResponse)
	resp, err := c.Do(ctx, req, restart)
	if err != nil {
		return nil, resp, err
	}

	return restart, resp, nil
}
