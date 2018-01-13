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
