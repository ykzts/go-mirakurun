// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// ProgramsService ...
type ProgramsService service

// Program ...
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

// ProgramGenre ...
type ProgramGenre struct {
	Level1      int `json:"lv1"`
	Level2      int `json:"lv2"`
	UserNibble1 int `json:"un1"`
	UserNibble2 int `json:"un2"`
}

// ProgramVideo ...
type ProgramVideo struct {
	Type       string `json:"type"`
	Resolution string `json:"resolution"`

	StreamContent int `json:"streamContent"`
	ComponentType int `json:"componentType"`
}

// ProgramAudio ...
type ProgramAudio struct {
	SamplingRate  int `json:"samplingRate"`
	ComponentType int `json:"componentType"`
}

// ProgramSeries ...
type ProgramSeries struct {
	ID          int       `json:"id"`
	Repeat      int       `json:"repeat"`
	Pattern     int       `json:"pattern"`
	ExpiresAt   Timestamp `json:"expiresAt"`
	Episode     int       `json:"episode"`
	LastEpisode int       `json:"lastEpisode"`
	Name        string    `json:"name"`
}

// ProgramRelatedItem ...
type ProgramRelatedItem struct {
	NetworkID int `json:"networkId"`
	ServiceID int `json:"serviceId"`
	EventID   int `json:"eventId"`
}

// Get ...
func (s *ProgramsService) Get(ctx context.Context, id int) (*Program, *http.Response, error) {
	u := fmt.Sprintf("programs/%d", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	program := new(Program)
	resp, err := s.client.Do(ctx, req, program)
	if err != nil {
		return nil, resp, err
	}

	return program, resp, nil
}

// GetStream ...
func (s *ProgramsService) GetStream(ctx context.Context, id int, decode bool, w io.Writer) (*http.Response, error) {
	opt := &DecodeOptions{Decode: 0}
	if decode {
		opt.Decode = 1
	}

	u, err := addOptions(fmt.Sprintf("programs/%d/stream", id), opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, w)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ProgramsListOptions ...
type ProgramsListOptions struct {
	NetworkID int `url:"networkId,omitempty"`
	ServiceID int `url:"serviceId,omitempty"`
	EventID   int `url:"eventId,omitempty"`
}

// List ...
func (s *ProgramsService) List(ctx context.Context, opt *ProgramsListOptions) ([]*Program, *http.Response, error) {
	u, err := addOptions("programs", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	programs := []*Program{}
	resp, err := s.client.Do(ctx, req, &programs)
	if err != nil {
		return nil, resp, err
	}

	return programs, resp, nil
}
