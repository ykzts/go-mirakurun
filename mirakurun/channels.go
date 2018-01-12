// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"fmt"
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

// ChannelsListOptions specifies the optional parameters to the Client.GetChannels method.
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
