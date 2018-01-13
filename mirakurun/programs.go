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
	"fmt"
	"io"
	"net/http"
)

// Program represents a Mirakurun program.
type Program struct {
	ID        int       `json:"id"`
	EventID   int       `json:"eventId"`
	ServiceID int       `json:"serviceId"`
	NetworkID int       `json:"networkId"`
	StartAt   Timestamp `json:"startAt"`
	Duration  int       `json:"duration"`
	IsFree    bool      `json:"isFree"`

	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Genres      []ProgramGenre `json:"genres,omitempty"`
	Video       ProgramVideo   `json:"video,omitempty"`
	Audio       ProgramAudio   `json:"audio,omitempty"`

	Series ProgramSeries `json:"series,omitempty"`

	Extended map[string]string `json:"extended,omitempty"`

	RelatedItems []ProgramRelatedItem `json:"relatedItems,omitempty"`
}

// ProgramGenre represents a Mirakurun program genre.
type ProgramGenre struct {
	Level1      int `json:"lv1"`
	Level2      int `json:"lv2"`
	UserNibble1 int `json:"un1"`
	UserNibble2 int `json:"un2"`
}

// ProgramVideo represents a Mirakurun program video.
type ProgramVideo struct {
	Type       string `json:"type"`
	Resolution string `json:"resolution"`

	StreamContent int `json:"streamContent"`
	ComponentType int `json:"componentType"`
}

// ProgramAudio represents a Mirakurun program audio.
type ProgramAudio struct {
	SamplingRate  int `json:"samplingRate"`
	ComponentType int `json:"componentType"`
}

// ProgramSeries represents a Mirakurun program series.
type ProgramSeries struct {
	ID          int       `json:"id"`
	Repeat      int       `json:"repeat"`
	Pattern     int       `json:"pattern"`
	ExpiresAt   Timestamp `json:"expiresAt"`
	Episode     int       `json:"episode"`
	LastEpisode int       `json:"lastEpisode"`
	Name        string    `json:"name"`
}

// ProgramRelatedItem represents a Mirakurun program related item.
type ProgramRelatedItem struct {
	NetworkID int `json:"networkId"`
	ServiceID int `json:"serviceId"`
	EventID   int `json:"eventId"`
}

// ProgramsListOptions ...
type ProgramsListOptions struct {
	NetworkID int `url:"networkId,omitempty"`
	ServiceID int `url:"serviceId,omitempty"`
	EventID   int `url:"eventId,omitempty"`
}

// GetPrograms lists the programs.
func (c *Client) GetPrograms(ctx context.Context, opt *ProgramsListOptions) ([]*Program, *http.Response, error) {
	u, err := addOptions("programs", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	programs := []*Program{}
	resp, err := c.Do(ctx, req, &programs)
	if err != nil {
		return nil, resp, err
	}

	return programs, resp, nil
}

// GetProgram fetches a program.
func (c *Client) GetProgram(ctx context.Context, id int) (*Program, *http.Response, error) {
	u := fmt.Sprintf("programs/%d", id)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	program := new(Program)
	resp, err := c.Do(ctx, req, program)
	if err != nil {
		return nil, resp, err
	}

	return program, resp, nil
}

// GetProgramStream fetches a program stream.
func (c *Client) GetProgramStream(ctx context.Context, id int, decode bool) (io.ReadCloser, *http.Response, error) {
	u := fmt.Sprintf("programs/%d/stream", id)

	return c.getTS(ctx, u, decode)
}
