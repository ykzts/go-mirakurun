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
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_GetChannels(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/channels", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/channels.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	channels, _, err := c.GetChannels(ctx, nil)
	if err != err {
		t.Errorf("Client.GetChannels returned error: %v", err)
	}

	if got, want := channels[0].Name, "TOKYO MX"; got != want {
		t.Errorf("Channel name is %v, want %v", got, want)
	}
}

func TestClient_GetChannelsByType(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/channels/GR", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/channels.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	channels, _, err := c.GetChannelsByType(ctx, "GR", nil)
	if err != err {
		t.Errorf("Client.GetChannels returned error: %v", err)
	}

	if got, want := channels[0].Name, "TOKYO MX"; got != want {
		t.Errorf("Channel name is %v, want %v", got, want)
	}
}

func TestClient_GetChannel(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/channels/GR/16", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/channel.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	channel, _, err := c.GetChannel(ctx, "GR", "16")
	if err != err {
		t.Errorf("Client.GetChannel returned error: %v", err)
	}

	if got, want := channel.Name, "TOKYO MX"; got != want {
		t.Errorf("Channel name is %v, want %v", got, want)
	}
}
