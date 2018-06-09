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

func TestClient_GetServicesByChannel(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/channels/GR/16/services", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/services?channel.type=GR&channel.channel=16", 307)
	})
	mux.HandleFunc("/api/services", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		if q.Get("channel.type") == "GR" && q.Get("channel.channel") == "16" {
			http.ServeFile(w, r, "testdata/services.json")
		} else {
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, "[]")
		}
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	services, _, err := c.GetServicesByChannel(context.Background(), "GR", "16")
	if err != nil {
		t.Fatal(err)
	}

	if len(services) < 1 {
		t.Errorf("services is empty")
	} else if got, want := services[0].Name, "TOKYO MX1"; got != want {
		t.Errorf("service name is %v, want %v", got, want)
	}
}

func TestClient_GetServiceByChannel(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/channels/GR/16/services/23608", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/services/3239123608", 307)
	})
	mux.HandleFunc("/api/services/3239123608", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/service.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	service, _, err := c.GetServiceByChannel(context.Background(), "GR", "16", 23608)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := service.Name, "TOKYO MX1"; got != want {
		t.Errorf("service name is %v, want %v", got, want)
	}
}

func TestClient_GetServices(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/services", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/services.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	services, _, err := c.GetServices(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(services) < 1 {
		t.Errorf("services is empty")
	} else if got, want := services[0].Name, "TOKYO MX1"; got != want {
		t.Errorf("service name is %v, want %v", got, want)
	}
}

func TestClient_GetService(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/services/3239123608", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/service.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	service, _, err := c.GetService(context.Background(), 3239123608)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := service.Name, "TOKYO MX1"; got != want {
		t.Errorf("service name is %v, want %v", got, want)
	}
}
