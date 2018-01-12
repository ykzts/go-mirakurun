// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"net/http"
)

// Status represents a Mirakurun status.
type Status struct {
	Version       string              `json:"version"`
	Process       ProcessStatus       `json:"process"`
	EPG           EPGStatus           `json:"epg"`
	StreamCount   map[string]int      `json:"streamCount"`
	ErrorCount    map[string]int      `json:"errorCount"`
	TimerAccuracy TimerAccuracyStatus `json:"timerAccuracy"`
}

// ProcessStatus represents a Mirakurun process status.
type ProcessStatus struct {
	Arch        string            `json:"arch"`
	Platform    string            `json:"platform"`
	Versions    map[string]string `json:"versions"`
	PID         int               `json:"pid"`
	MemoryUsage map[string]int    `json:"memoryUsage"`
}

// EPGStatus represents a Mirakurun EPG status.
type EPGStatus struct {
	GatheringNetworks []int `json:"gatheringNetworks"`
	StoredEvents      int   `json:"storedEvents"`
}

// TimerAccuracyStatus represents a Mirakurun time accuracy status.
type TimerAccuracyStatus struct {
	Last float64            `json:"last"`
	M1   map[string]float64 `json:"m1"`
	M5   map[string]float64 `json:"m5"`
	M15  map[string]float64 `json:"m15"`
}

// GetStatus fetches a status.
func (c *Client) GetStatus(ctx context.Context) (*Status, *http.Response, error) {
	req, err := c.NewRequest("GET", "status", nil)
	if err != nil {
		return nil, nil, err
	}

	status := new(Status)
	resp, err := c.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
