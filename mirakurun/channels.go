package mirakurun

import (
	"context"
	"fmt"
	"net/http"
)

// ChannelsService ...
type ChannelsService service

// Channel represents a Mirakurun channel.
type Channel struct {
	Type     string    `json:"type"`
	Channel  string    `json:"channel"`
	Name     string    `json:"name,omitempty"`
	Satelite string    `json:"satelite,omitempty"`
	Space    int64     `json:"space,omitempty"`
	Services []Service `json:"services,omitempty"`
}

// Get ...
func (s *ChannelsService) Get(ctx context.Context, typ string, channelID string) (*Channel, *http.Response, error) {
	u := fmt.Sprintf("channels/%s/%s", typ, channelID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channel := new(Channel)
	resp, err := s.client.Do(ctx, req, channel)
	if err != nil {
		return nil, resp, err
	}

	return channel, resp, nil
}

// ChannelsListOptions specifies optinal parameters to the ChannelsService.List method.
type ChannelsListOptions struct {
	Type    string `url:"type,omitempty"`
	Channel string `url:"channel,omitempty"`
	Name    string `url:"name,omitempty"`
}

// List ...
func (s *ChannelsService) List(ctx context.Context, opt *ChannelsListOptions) ([]*Channel, *http.Response, error) {
	u, err := addOptions("channels", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channels := []*Channel{}
	resp, err := s.client.Do(ctx, req, &channels)
	if err != nil {
		return nil, resp, err
	}

	return channels, resp, nil
}

// ListByType ...
func (s *ChannelsService) ListByType(ctx context.Context, typ string, opt *ChannelsListOptions) ([]*Channel, *http.Response, error) {
	u, err := addOptions(fmt.Sprintf("channels/%s", typ), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channels := []*Channel{}
	resp, err := s.client.Do(ctx, req, &channels)
	if err != nil {
		return nil, resp, err
	}

	return channels, resp, nil
}
