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

func TestNewClient(t *testing.T) {
	c := NewClient()

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
}

func TestClient_NewRequest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/channels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "[]")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	req, err := c.NewRequest("GET", "channels", nil)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := req.URL.String(), server.URL+"/api/channels"; got != want {
		t.Errorf("request URL is %v, want %v", got, want)
	}
}

func TestClient_Do(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/channels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "[]")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	req, err := c.NewRequest("GET", "channels", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Do(context.Background(), req, nil)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := resp.StatusCode, 200; got != want {
		t.Errorf("status code is %v, want %v", got, want)
	}
}

func TestClient_Do_notFound(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/not-found", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Cannot GET /api/not-found\n", http.StatusNotFound)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	req, err := c.NewRequest("GET", "not-found", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Do(context.Background(), req, nil)
	if err == nil {
		t.Fatal("request should returns error")
	}

	if got, want := err.Error(), "mirakurun: 404 Not Found"; got != want {
		t.Errorf("error is %v, want %v", got, want)
	}
}
