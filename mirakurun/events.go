// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"io"
	"net/http"
)

// EventsService ...
type EventsService service

// Event ...
type Event struct {
	Resource string      `json:"resource"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
	Time     Timestamp   `json:"time"`
}

// List ...
func (s *EventsService) List(ctx context.Context) ([]*Event, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "events", nil)
	if err != nil {
		return nil, nil, err
	}

	events := []*Event{}
	resp, err := s.client.Do(ctx, req, &events)
	if err != nil {
		return nil, resp, err
	}

	return events, resp, nil
}

// EventsListOptions ...
type EventsListOptions struct {
	Resource string `url:"resource,omitempty"`
	Type     string `url:"type,omitempty"`
}

// ListStream ...
func (s *EventsService) ListStream(ctx context.Context, opt *EventsListOptions, w io.Writer) (*http.Response, error) {
	u, err := addOptions("events/stream", opt)
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
