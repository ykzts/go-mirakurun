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

func TestClient_GetTuners(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tuners", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/tuners.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	tuners, _, err := c.GetTuners(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(tuners) < 1 {
		t.Error("tuners is empty")
	} else if got, want := tuners[0].Name, "PT3-T1"; got != want {
		t.Errorf("tuner name is %v, want %v", got, want)
	}
}

func TestClient_GetTuner(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tuners/0", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/tuner.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	tuner, _, err := c.GetTuner(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := tuner.Name, "PT3-T1"; got != want {
		t.Errorf("tuner name is %v, want %v", got, want)
	}
}

func TestClient_GetTunerProcess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tuners/0/process", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/tuner_process.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	p, _, err := c.GetTunerProcess(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := p.PID, 1000; got != want {
		t.Errorf("tuner process is %v, want %v", got, want)
	}
}

func TestClient_KillTunerProcess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tuners/0/process", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			http.ServeFile(w, r, "testdata/tuner_process.json")
		} else {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	p, _, err := c.KillTunerProcess(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := p.PID, 1000; got != want {
		t.Errorf("tuner process is %v, want %v", got, want)
	}
}

func TestClient_GetTunersConfig(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config/tuners", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/tuners_config.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	tuners, _, err := c.GetTunersConfig(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(tuners) < 1 {
		t.Error("tuners is empty")
	} else if got, want := tuners[0].Name, "PT3-T1"; got != want {
		t.Errorf("tuner name is %v, want %v", got, want)
	}
}

func TestClient_UpdateTunersConfig(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/config/tuners", func(w http.ResponseWriter, r *http.Request) {
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

	body := TunersConfig{
		&TunerConfig{
			Name:       "PT3-T1",
			Types:      []string{"GR"},
			Command:    "recpt1 --device /dev/pt3video2 <channel> - -",
			Decoder:    "arib-b25-stream-test",
			IsDisabled: false,
		},
	}
	tuners, _, err := c.UpdateTunersConfig(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if len(tuners) < 1 {
		t.Error("tuners is empty")
	} else if got, want := tuners[0].Name, "PT3-T1"; got != want {
		t.Errorf("tuner name is %v, want %v", got, want)
	}
}
