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

// TunersConfig represents a Mirakurun tuners config.
type TunersConfig []*TunerConfig

// TunerConfig represents a Mirakurun tuner config.
type TunerConfig struct {
	Name          string   `json:"name"`
	Types         []string `json:"types"`
	Command       string   `json:"command,omitempty"`
	DVBDevicePath string   `json:"dvbDevicePath,omitempty"`
	Decoder       string   `json:"decoder,omitempty"`
	IsDisabled    bool     `json:"isDisabled,omitempty"`
}

// GetTunersConfig fetches a tuners config.
func (c *Client) GetTunersConfig(ctx context.Context) (TunersConfig, *http.Response, error) {
	req, err := c.NewRequest("GET", "config/tuners", nil)
	if err != nil {
		return nil, nil, err
	}

	var tuners TunersConfig
	resp, err := c.Do(ctx, req, &tuners)
	if err != nil {
		return nil, resp, err
	}

	return tuners, resp, nil
}

// UpdateTunersConfig updates a tuners config.
func (c *Client) UpdateTunersConfig(ctx context.Context, body TunersConfig) (TunersConfig, *http.Response, error) {
	req, err := c.NewRequest("PUT", "config/tuners", body)
	if err != nil {
		return nil, nil, err
	}

	var tuners TunersConfig
	resp, err := c.Do(ctx, req, &tuners)
	if err != nil {
		return nil, resp, err
	}

	return tuners, resp, nil
}
