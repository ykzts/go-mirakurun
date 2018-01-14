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

func TestClient_GetPrograms(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/programs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/programs.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	programs, _, err := c.GetPrograms(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(programs) < 1 {
		t.Error("programs is empty")
	} else if got, want := programs[0].Name, "Cardcaptor Sakura: Clear Card ep. 3"; got != want {
		t.Errorf("program name is %v, want %v", got, want)
	}
}

func TestClient_GetProgram(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/programs/40010310979", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/program.json")
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient()
	c.BaseURL, _ = url.Parse(server.URL + "/api/")

	program, _, err := c.GetProgram(context.Background(), 40010310979)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := program.Name, "Cardcaptor Sakura: Clear Card ep. 3"; got != want {
		t.Errorf("program name is %v, want %v", got, want)
	}
}
