// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// ServicesService ...
type ServicesService service

// Service represents a Mirakurun service.
type Service struct {
	ID                 int     `json:"id"`
	ServiceID          int     `json:"serviceId"`
	NetworkID          int     `json:"networkId"`
	Name               string  `json:"name"`
	LogoID             int     `json:"logoId,omitempty"`
	HasLogoData        bool    `json:"hasLogoData,omitempty"`
	RemoteControlKeyID int     `json:"remoteControlKeyId,omitempty"`
	Channel            Channel `json:"channel,omitempty"`
}

// Get ...
func (s *ServicesService) Get(ctx context.Context, id int) (*Service, *http.Response, error) {
	u := fmt.Sprintf("services/%d", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	service := new(Service)
	resp, err := s.client.Do(ctx, req, service)
	if err != nil {
		return nil, resp, err
	}

	return service, resp, nil
}

// GetStream ...
func (s *ServicesService) GetStream(ctx context.Context, id int, decode bool, w io.Writer) (*http.Response, error) {
	opt := &DecodeOptions{Decode: 0}
	if decode {
		opt.Decode = 1
	}

	u, err := addOptions(fmt.Sprintf("services/%d/stream", id), opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetByChannel ...
func (s *ServicesService) GetByChannel(ctx context.Context, typ string, channel string, sid int) (*Service, *http.Response, error) {
	u := fmt.Sprintf("channels/%s/%s/services/%d", typ, channel, sid)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	service := new(Service)
	resp, err := s.client.Do(ctx, req, service)
	if err != nil {
		return nil, resp, err
	}

	return service, resp, nil
}

// GetStreamByChannel ...
func (s *ServicesService) GetStreamByChannel(ctx context.Context, typ string, channel string, sid int, decode bool, w io.Writer) (*http.Response, error) {
	opt := &DecodeOptions{Decode: 0}
	if decode {
		opt.Decode = 1
	}

	u, err := addOptions(fmt.Sprintf("channels/%s/%s/services/%d/stream", typ, channel, sid), opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetLogo ...
func (s *ServicesService) GetLogo(ctx context.Context, id int, w io.Writer) (*http.Response, error) {
	u := fmt.Sprintf("services/%d/logo", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ServicesListOptions ...
type ServicesListOptions struct {
	ServiceID   int    `url:"serviceId,omitempty"`
	NetworkID   int    `url:"networkId,omitempty"`
	Name        string `url:"name,omitempty"`
	Type        int    `url:"type,omitempty"`
	ChannelType string `url:"channel.type,omitempty"`
	Channel     string `url:"channel.channel,omitempty"`
}

// List ...
func (s *ServicesService) List(ctx context.Context, opt *ServicesListOptions) ([]*Service, *http.Response, error) {
	u, err := addOptions("services", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	services := []*Service{}
	resp, err := s.client.Do(ctx, req, &services)
	if err != nil {
		return nil, resp, err
	}

	return services, resp, nil
}

// ListByChannel ...
func (s *ServicesService) ListByChannel(ctx context.Context, typ string, channel string) ([]*Service, *http.Response, error) {
	u := fmt.Sprintf("channels/%s/%s/services", typ, channel)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	services := []*Service{}
	resp, err := s.client.Do(ctx, req, &services)
	if err != nil {
		return nil, resp, err
	}

	return services, resp, nil
}
