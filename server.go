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
	"bytes"
	"context"
	"io"
	"net/http"
)

// ServerConfig represents a Mirakurun server config.
type ServerConfig struct {
	Path                      string `json:"path,omitempty"`
	Port                      int    `json:"port,omitempty"`
	DisableIPv6               bool   `json:"disableIPv6,omitempty"`
	LogLevel                  int    `json:"logLevel,omitempty"`
	MaxLogHistory             int    `json:"maxLogHistory,omitempty"`
	HighWaterMark             int    `json:"highWaterMark,omitempty"`
	OverflowTimeLimit         int    `json:"overflowTimeLimit,omitempty"`
	MaxBufferBytesBeforeReady int    `json:"maxBufferBytesBeforeReady,omitempty"`
	EventEndTimeout           int    `json:"eventEndTimeout,omitempty"`
	ProgramGCInterval         int    `json:"programGCInterval,omitempty"`
	EPGGatheringInterval      int    `json:"epgGatheringInterval,omitempty"`
	EPGRetrievalTime          int    `json:"epgRetrievalTime,omitempty"`
}

// GetServerConfig fetches a server config.
func (c *Client) GetServerConfig(ctx context.Context) (*ServerConfig, *http.Response, error) {
	req, err := c.NewRequest("GET", "config/server", nil)
	if err != nil {
		return nil, nil, err
	}

	config := new(ServerConfig)
	resp, err := c.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

// UpdateServerConfig updates a server config.
func (c *Client) UpdateServerConfig(ctx context.Context, body *ServerConfig) (*ServerConfig, *http.Response, error) {
	req, err := c.NewRequest("PUT", "config/server", body)
	if err != nil {
		return nil, nil, err
	}

	config := new(ServerConfig)
	resp, err := c.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

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
	return c.requestStream(ctx, "PUT", "log/stream")
}

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
