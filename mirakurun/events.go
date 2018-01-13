// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"io"
	"net/http"
)

// Event represents a Mirakurun event.
type Event struct {
	Resource string      `json:"resource"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
	Time     Timestamp   `json:"time"`
}

// EventsListOptions specifies the optional parameters to the Client.GetEventsStream method.
type EventsListOptions struct {
	Resource string `url:"resource,omitempty"`
	Type     string `url:"type,omitempty"`
}

// GetEvents lists the events.
func (c *Client) GetEvents(ctx context.Context) ([]*Event, *http.Response, error) {
	req, err := c.NewRequest("GET", "events", nil)
	if err != nil {
		return nil, nil, err
	}

	events := []*Event{}
	resp, err := c.Do(ctx, req, &events)
	if err != nil {
		return nil, resp, err
	}

	return events, resp, nil
}

// GetEventsStream fetches a events stream.
func (c *Client) GetEventsStream(ctx context.Context, opt *EventsListOptions) (io.ReadCloser, *http.Response, error) {
	u, err := addOptions("events/stream", opt)
	if err != nil {
		return nil, nil, err
	}

	return c.requestStream(ctx, "GET", u)
}
