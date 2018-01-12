// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

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
