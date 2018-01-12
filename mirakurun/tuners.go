// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"fmt"
	"net/http"
)

// TunerDevice represents a Mirakurun tuner device.
type TunerDevice struct {
	Index       int         `json:"index"`
	Name        string      `json:"name"`
	Types       []string    `json:"types"`
	Command     string      `json:"command"`
	PID         int         `json:"pid"`
	Users       []TunerUser `json:"users"`
	IsAvailable bool        `json:"isAvailable"`
	IsFree      bool        `json:"isFree"`
	IsUsing     bool        `json:"isUsing"`
	IsFault     bool        `json:"isFault"`
}

// TunerUser represents a Mirakurun tuner user.
type TunerUser struct {
	ID       string `json:"id"`
	Priority int    `json:"priority"`
	Agent    string `json:"agent,omitempty"`
}

// TunerProcess represents a Mirakurun tuner process.
type TunerProcess struct {
	PID int `json:"pid"`
}

// GetTuners lists the tuners.
func (c *Client) GetTuners(ctx context.Context) ([]*TunerDevice, *http.Response, error) {
	req, err := c.NewRequest("GET", "tuners", nil)
	if err != nil {
		return nil, nil, err
	}

	tuners := []*TunerDevice{}
	resp, err := c.Do(ctx, req, &tuners)
	if err != nil {
		return nil, resp, err
	}

	return tuners, resp, nil
}

// GetTuner fetches a tuner.
func (c *Client) GetTuner(ctx context.Context, index int) (*TunerDevice, *http.Response, error) {
	u := fmt.Sprintf("tuners/%d", index)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tuner := new(TunerDevice)
	resp, err := c.Do(ctx, req, tuner)
	if err != nil {
		return nil, resp, err
	}

	return tuner, resp, nil
}

// GetTunerProcess fetches a tuner process.
func (c *Client) GetTunerProcess(ctx context.Context, index int) (*TunerProcess, *http.Response, error) {
	u := fmt.Sprintf("tuners/%d/process", index)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	process := new(TunerProcess)
	resp, err := c.Do(ctx, req, process)
	if err != nil {
		return nil, resp, err
	}

	return process, resp, nil
}

// KillTunerProcess kills a tuner process.
func (c *Client) KillTunerProcess(ctx context.Context, index int) (*TunerProcess, *http.Response, error) {
	u := fmt.Sprintf("tuners/%d/process", index)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	process := new(TunerProcess)
	resp, err := c.Do(ctx, req, process)
	if err != nil {
		return nil, resp, err
	}

	return process, resp, nil
}
