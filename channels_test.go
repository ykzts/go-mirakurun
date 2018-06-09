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

	channels, _, err := c.GetChannels(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(channels) < 1 {
		t.Errorf("channels is empty")
	}
	if got, want := channels[0].Name, "TOKYO MX"; got != want {
		t.Errorf("channel name is %v, want %v", got, want)
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

	channels, _, err := c.GetChannelsByType(context.Background(), "GR", nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(channels) < 1 {
		t.Errorf("channels is empty")
	} else if got, want := channels[0].Name, "TOKYO MX"; got != want {
		t.Errorf("channel name is %v, want %v", got, want)
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

	channel, _, err := c.GetChannel(context.Background(), "GR", "16")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := channel.Name, "TOKYO MX"; got != want {
		t.Errorf("channel name is %v, want %v", got, want)
	}
}

func TestClient_GetChannelsConfig(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config/channels", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/channels_config.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	channels, _, err := c.GetChannelsConfig(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(channels) < 1 {
		t.Error("channels is empty")
	} else if got, want := channels[0].Name, "TOKYO MX"; got != want {
		t.Errorf("channel name is %v, want %v", got, want)
	}
}

func TestClient_UpdateChannelsConfig(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config/channels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			w.Header().Set("Content-Type", "application/json")
			io.Copy(w, r.Body)
		} else {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	body := ChannelsConfig{
		&ChannelConfig{
			Name:    "TOKYO MX",
			Type:    "GR",
			Channel: "16",
		},
	}
	channels, _, err := c.UpdateChannelsConfig(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if len(channels) < 1 {
		t.Error("channels is empty")
	} else if got, want := channels[0].Name, "TOKYO MX"; got != want {
		t.Errorf("channel name is %v, want %v", got, want)
	}
}

func TestClient_ChannelScan(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config/channels/scan", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
			io.WriteString(w, "channel: scanning... (type: \"GR\")\n\n")
		} else {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	stream, resp, err := c.ChannelScan(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("status code is %v, want %v", got, want)
	}
}
