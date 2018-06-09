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

func TestClient_GetServerConfig(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config/server", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/server_config.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	serverConfig, _, err := c.GetServerConfig(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := serverConfig.Port, 40772; got != want {
		t.Errorf("server port is %v, want %v", got, want)
	}
}

func TestClient_UpdateServerConfig(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config/server", func(w http.ResponseWriter, r *http.Request) {
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

	body := &ServerConfig{
		Path:     "/var/run/mirakurun.sock",
		Port:     40772,
		LogLevel: 2,
	}
	serverConfig, _, err := c.UpdateServerConfig(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := serverConfig.Port, 40772; got != want {
		t.Errorf("server port is %v, want %v", got, want)
	}
}

func TestClient_GetLog(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/log", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/log.txt")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	_, resp, err := c.GetLog(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := resp.Header.Get("Content-Type"), "text/plain; charset=utf-8"; got != want {
		t.Errorf("content type is %v, want %v", got, want)
	}
}

func TestClient_GetLogStream(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/log/stream", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	stream, resp, err := c.GetLogStream(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("status code is %v, want %v", got, want)
	}
}

func TestClient_CheckVersion(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/version.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	version, _, err := c.CheckVersion(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := version.Latest, "2.5.7"; got != want {
		t.Errorf("version is %v, want %v", got, want)
	}
}

func TestClient_UpdateVersion(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/version/update", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, `Updating...`)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	stream, resp, err := c.UpdateVersion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	if got, want := resp.StatusCode, http.StatusAccepted; got != want {
		t.Errorf("status code is %v, want %v", got, want)
	}
}

func TestClient_GetStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"version":"2.0.0"}`)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	status, _, err := c.GetStatus(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := status.Version, "2.0.0"; got != want {
		t.Errorf("version is %v, want %v", got, want)
	}
}

func TestClient_Restart(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/restart", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			io.WriteString(w, `{"_cmd_pid":1000}`)
		} else {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	restart, _, err := c.Restart(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := restart.PID, 1000; got != want {
		t.Errorf("pid is %v, want %v", got, want)
	}
}
