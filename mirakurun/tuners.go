// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"fmt"
	"net/http"
)

// TunersService ...
type TunersService service

// TunerDevice ...
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

// TunerUser ...
type TunerUser struct {
	ID       string `json:"id"`
	Priority int    `json:"priority"`
	Agent    string `json:"agent,omitempty"`
}

// TunerProcess ..
type TunerProcess struct {
	PID int `json:"pid"`
}

// Get ...
func (s *TunersService) Get(ctx context.Context, index int) (*TunerDevice, *http.Response, error) {
	u := fmt.Sprintf("tuners/%d", index)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tuner := new(TunerDevice)
	resp, err := s.client.Do(ctx, req, tuner)
	if err != nil {
		return nil, resp, err
	}

	return tuner, resp, nil
}

// GetProcess
func (s *TunersService) GetProcess(ctx context.Context, index int) (*TunerProcess, *http.Response, error) {
	u := fmt.Sprintf("tuners/%d/process", index)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	process := new(TunerProcess)
	resp, err := s.client.Do(ctx, req, process)
	if err != nil {
		return nil, resp, err
	}

	return process, resp, nil
}

// KillProcess
func (s *TunersService) KillProcess(ctx context.Context, index int) (*TunerProcess, *http.Response, error) {
	u := fmt.Sprintf("tuners/%d/process", index)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	process := new(TunerProcess)
	resp, err := s.client.Do(ctx, req, process)
	if err != nil {
		return nil, resp, err
	}

	return process, resp, nil
}

// List ...
func (s *TunersService) List(ctx context.Context) ([]*TunerDevice, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "tuners", nil)
	if err != nil {
		return nil, nil, err
	}

	tuners := []*TunerDevice{}
	resp, err := s.client.Do(ctx, req, &tuners)
	if err != nil {
		return nil, resp, err
	}

	return tuners, resp, nil
}
