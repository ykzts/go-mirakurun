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

// ServicesListOptions specifies the optional parameters to the Client.GetServices method.
type ServicesListOptions struct {
	ServiceID   int    `url:"serviceId,omitempty"`
	NetworkID   int    `url:"networkId,omitempty"`
	Name        string `url:"name,omitempty"`
	Type        int    `url:"type,omitempty"`
	ChannelType string `url:"channel.type,omitempty"`
	Channel     string `url:"channel.channel,omitempty"`
}

// GetServicesByChannel lists the services for the specified channel.
func (c *Client) GetServicesByChannel(ctx context.Context, typ string, channel string) ([]*Service, *http.Response, error) {
	u := fmt.Sprintf("channels/%s/%s/services", typ, channel)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	services := []*Service{}
	resp, err := c.Do(ctx, req, &services)
	if err != nil {
		return nil, resp, err
	}

	return services, resp, nil
}

// GetServiceByChannel fetches a service for the specified channel.
func (c *Client) GetServiceByChannel(ctx context.Context, typ string, channel string, sid int) (*Service, *http.Response, error) {
	u := fmt.Sprintf("channels/%s/%s/services/%d", typ, channel, sid)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	service := new(Service)
	resp, err := c.Do(ctx, req, service)
	if err != nil {
		return nil, resp, err
	}

	return service, resp, nil
}

// GetServiceStreamByChannel fetches a service stream for the specified channel.
func (c *Client) GetServiceStreamByChannel(ctx context.Context, typ string, channel string, sid int, decode bool) (io.ReadCloser, *http.Response, error) {
	opt := &DecodeOptions{Decode: 0}
	if decode {
		opt.Decode = 1
	}

	u, err := addOptions(fmt.Sprintf("channels/%s/%s/services/%d/stream", typ, channel, sid), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
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

// GetServices lists the services.
func (c *Client) GetServices(ctx context.Context, opt *ServicesListOptions) ([]*Service, *http.Response, error) {
	u, err := addOptions("services", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	services := []*Service{}
	resp, err := c.Do(ctx, req, &services)
	if err != nil {
		return nil, resp, err
	}

	return services, resp, nil
}

// GetService fetches a service.
func (c *Client) GetService(ctx context.Context, id int) (*Service, *http.Response, error) {
	u := fmt.Sprintf("services/%d", id)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	service := new(Service)
	resp, err := c.Do(ctx, req, service)
	if err != nil {
		return nil, resp, err
	}

	return service, resp, nil
}

// GetLogoImage fetches a service logo stream.
func (c *Client) GetLogoImage(ctx context.Context, id int) (io.ReadCloser, *http.Response, error) {
	u := fmt.Sprintf("services/%d/logo", id)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	return resp.Body, resp, nil
}

// GetServiceStream fetches a service stream.
func (c *Client) GetServiceStream(ctx context.Context, id int, decode bool) (io.ReadCloser, *http.Response, error) {
	opt := &DecodeOptions{Decode: 0}
	if decode {
		opt.Decode = 1
	}

	u, err := addOptions(fmt.Sprintf("services/%d/stream", id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
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
