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
