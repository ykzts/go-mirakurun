// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"io"
	"net/http"
)

// ChannelsConfig represents a Mirakurun channels config.
type ChannelsConfig []*ChannelConfig

// ChannelConfig represents a Mirakurun channel config
type ChannelConfig struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Channel    string `json:"channel"`
	Satelite   string `json:"satelite,omitempty"`
	ServiceID  int    `json:"serviceId,omitempty"`
	Space      int    `json:"space,omitempty"`
	IsDisabled bool   `json:"isDisabled,omitempty"`
}

// GetChannelsConfig fetches a channels config.
func (c *Client) GetChannelsConfig(ctx context.Context) (ChannelsConfig, *http.Response, error) {
	req, err := c.NewRequest("GET", "config/channels", nil)
	if err != nil {
		return nil, nil, err
	}

	var config ChannelsConfig
	resp, err := c.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

// UpdateChannelsConfig updates a channels config.
func (c *Client) UpdateChannelsConfig(ctx context.Context, body ChannelsConfig) (ChannelsConfig, *http.Response, error) {
	req, err := c.NewRequest("PUT", "config/channels", body)
	if err != nil {
		return nil, nil, err
	}

	var config ChannelsConfig
	resp, err := c.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

// ChannelScanOptions specifies the optional parameters to the Client.ChannelScan method.
type ChannelScanOptions struct {
	Type string `url:"type,omitempty"`
	Min  int    `url:"min,omitempty"`
	Max  int    `url:"max,omitempty"`
}

// ChannelScan scans a channels.
func (c *Client) ChannelScan(ctx context.Context, opt *ChannelScanOptions) (io.ReadCloser, *http.Response, error) {
	u, err := addOptions("config/channels/scan", opt)
	if err != nil {
		return nil, nil, err
	}

	return c.requestStream(ctx, "PUT", u)
}

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

	var config TunersConfig
	resp, err := c.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

// UpdateTunersConfig updates a tuners config.
func (c *Client) UpdateTunersConfig(ctx context.Context, body TunersConfig) (TunersConfig, *http.Response, error) {
	req, err := c.NewRequest("PUT", "config/tuners", body)
	if err != nil {
		return nil, nil, err
	}

	var config TunersConfig
	resp, err := c.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}
