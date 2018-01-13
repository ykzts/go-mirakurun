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
	"io"
	"net/http"
)

// Channel represents a Mirakurun channel.
type Channel struct {
	Type     string    `json:"type"`
	Channel  string    `json:"channel"`
	Name     string    `json:"name,omitempty"`
	Satelite string    `json:"satelite,omitempty"`
	Space    int64     `json:"space,omitempty"`
	Services []Service `json:"services,omitempty"`
}

// ChannelsListOptions specifies the optional parameters to the Client.GetChannels method and Client.GetChannelsByType method.
type ChannelsListOptions struct {
	Type    string `url:"type,omitempty"`
	Channel string `url:"channel,omitempty"`
	Name    string `url:"name,omitempty"`
}

// GetChannels lists the channels.
func (c *Client) GetChannels(ctx context.Context, opt *ChannelsListOptions) ([]*Channel, *http.Response, error) {
	u, err := addOptions("channels", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channels := []*Channel{}
	resp, err := c.Do(ctx, req, &channels)
	if err != nil {
		return nil, resp, err
	}

	return channels, resp, nil
}

// GetChannelsByType lists the channels for the specified type.
func (c *Client) GetChannelsByType(ctx context.Context, typ string, opt *ChannelsListOptions) ([]*Channel, *http.Response, error) {
	u, err := addOptions(fmt.Sprintf("channels/%s", typ), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channels := []*Channel{}
	resp, err := c.Do(ctx, req, &channels)
	if err != nil {
		return nil, resp, err
	}

	return channels, resp, nil
}

// GetChannel fetches a channel.
func (c *Client) GetChannel(ctx context.Context, typ string, channelID string) (*Channel, *http.Response, error) {
	u := fmt.Sprintf("channels/%s/%s", typ, channelID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channel := new(Channel)
	resp, err := c.Do(ctx, req, channel)
	if err != nil {
		return nil, resp, err
	}

	return channel, resp, nil
}

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
