package mirakurun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBasePath  = "/api"
	defaultUserAgent = "go-mirakurun/1.0.0"
)

// Channel ...
type Channel struct {
	Type     string    `json:"type"`
	Channel  string    `json:"channel"`
	Name     string    `json:"name,omitempty"`
	Satelite string    `json:"satelite,omitempty"`
	Space    int64     `json:"space,omitempty"`
	Services []Service `json:"services,omitempty"`
}

// Service ...
type Service struct {
	ID                 int64   `json:"id"`
	ServiceID          int64   `json:"serviceId"`
	NetworkID          int64   `json:"networkId"`
	Name               string  `json:"name"`
	LogoID             int64   `json:"logoId,omitempty"`
	HasLogoData        bool    `json:"hasLogoData,omitempty"`
	RemoteControlKeyID int64   `json:"remoteControlKeyId,omitempty"`
	Channel            Channel `json:"channel,omitempty"`
}

// Program

// ProgramGenre

// ProgramSeries

// ProgramRelatedItem

// TunerDevice

// TunerUser

// TunerProcess

// Event

// ConfigServer

// ConfigTunersItem

// ConfigChannelsItem

// Version ...
type Version struct {
	Current string `json:"current"`
	Latest  string `json:"latest"`
}

// Status ...
type Status struct {
	Version       string              `json:"version"`
	Process       ProcessStatus       `json:"process"`
	EPG           EPGStatus           `json:"epg"`
	StreamCount   map[string]int64    `json:"streamCount"`
	ErrorCount    map[string]int64    `json:"errorCount"`
	TimerAccuracy TimerAccuracyStatus `json:"timerAccuracy"`
}

// ProcessStatus ...
type ProcessStatus struct {
	Arch        string            `json:"arch"`
	Platform    string            `json:"platform"`
	Versions    map[string]string `json:"versions"`
	PID         int64             `json:"pid"`
	MemoryUsage map[string]int64  `json:"memoryUsage"`
}

// EPGStatus ...
type EPGStatus struct {
	GatheringNetworks []int64 `json:"gatheringNetworks"`
	StoredEvents      int64   `json:"storedEvents"`
}

// TimerAccuracyStatus ...
type TimerAccuracyStatus struct {
	Last float64            `json:"last"`
	M1   map[string]float64 `json:"m1"`
	M5   map[string]float64 `json:"m5"`
	M15  map[string]float64 `json:"m15"`
}

// RequestOption ...
type RequestOption struct {
	Priority int
	Headers  http.Header
	Query    interface{}
	Body     interface{}
}

// Response ...
type Response struct {
	Status      int
	StatusText  string
	ContentType string
	Headers     http.Header
	IsSuccess   bool
	Body        *bytes.Buffer
}

// RestartResponse ...
type RestartResponse struct {
	CmdPID int64 `json:"_cmd_pid"`
}

// ChannelsQuery ...
type ChannelsQuery struct {
	Type    string `url:"type,omitempty"`
	Channel string `url:"channel,omitempty"`
	Name    string `url:"name,omitempty"`
}

// ProgramsQuery ...
type ProgramsQuery struct {
	NetworkID int64 `url:"networkId,omitempty"`
	ServiceID int64 `url:"serviceId,omitempty"`
	EventID   int64 `url:"eventId,omitempty"`
}

// EventsQuery ...
type EventsQuery struct {
	Resource string `url:"resource,omitempty"`
	Type     string `url:"type,omitempty"`
}

// ServicesQuery ...
type ServicesQuery struct {
	ServiceID   int64  `url:"serviceId,omitempty"`
	NetworkID   int64  `url:"networkId,omitempty"`
	Name        string `url:"name,omitempty"`
	Type        int64  `url:"type,omitempty"`
	ChannelType string `url:"channel.type,omitempty"`
	Channel     string `url:"channel.channel,omitempty"`
}

// UpdateQuery ...
type UpdateQuery struct {
	Force bool `json:"force"`
}

// DecodableQuery ...
type DecodableQuery struct {
	Decode int `url:"decode,omitempty"`
}

// ChannelScanOption ...
type ChannelScanOption struct {
	Type string
	Min  int64
	Max  int64
}

// A Client manages communication with the Mirakurun API.
type Client struct {
	client *http.Client

	BasePath  string
	Priority  int
	Host      string
	Port      int
	UserAgent string
}

// NewClient returns a new Mirakurun API client.
func NewClient() *Client {
	httpClient := &http.Client{}

	c := &Client{client: httpClient, BasePath: defaultBasePath}

	return c
}

// Request ...
func (c *Client) Request(method, path string, option *RequestOption) (*Response, error) {
	resp, err := c.httpRequest(method, path, option)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contentType := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	ret := &Response{
		Status:      resp.StatusCode,
		StatusText:  resp.Status,
		ContentType: contentType,
		Headers:     resp.Header,
		IsSuccess:   resp.StatusCode >= 200 && resp.StatusCode <= 202,
		Body:        bytes.NewBuffer(nil),
	}

	io.Copy(ret.Body, resp.Body)

	return ret, nil
}

// GetChannels ...
func (c *Client) GetChannels(query ChannelsQuery) ([]Channel, error) {
	resp, err := c.Request("GET", "/channels", &RequestOption{Query: query})
	if err != nil {
		return nil, err
	}

	var channels []Channel
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&channels)

	return channels, nil
}

// GetChannelsByType ...
func (c *Client) GetChannelsByType(typ string, query ChannelsQuery) ([]Channel, error) {
	resp, err := c.Request("GET", "/channels/"+typ, &RequestOption{Query: query})
	if err != nil {
		return nil, err
	}

	var channels []Channel
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&channels)

	return channels, nil
}

// GetChannel

// GetServicesByChannel

// GetServiceByChannel

// GetServiceStreamByChannel

// GetServicesStreamByChannel

// GetChannelStream

// GetPrograms

// GetProgram

// GetProgramStream

// GetServices

// GetService

// GetLogoImage ...
func (c *Client) GetLogoImage(id int) (*bytes.Buffer, error) {
	resp, err := c.Request("GET", fmt.Sprintf("/services/%d/logo", id), &RequestOption{})
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// GetServiceStream

// GetTuners

// GetTuner

// GetTunerProcess

// KillTunerProcess

// GetEvents

// GetEventStream

// GetChannelsConfig

// UpdateChannelsConfig

// ChannelScan

// GetServerConfig

// UpdateServerConfig

// GetTunersConfig

// UpdateTunersConfig

// GetLog

// GetLogStream

// CheckVersion ...
func (c *Client) CheckVersion() (*Version, error) {
	resp, err := c.Request("GET", "/version", &RequestOption{})
	if err != nil {
		return nil, err
	}

	var version *Version
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&version)

	return version, nil
}

// UpdateVersion ...
func (c *Client) UpdateVersion(force bool) (*http.Response, error) {
	query := &UpdateQuery{Force: force}
	resp, err := c.requestStream("PUT", "/version/update", &RequestOption{Query: query})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetStatus ...
func (c *Client) GetStatus() (*Status, error) {
	resp, err := c.Request("GET", "/status", &RequestOption{})
	if err != nil {
		return nil, err
	}

	var status *Status
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&status)

	return status, nil
}

// Restart ...
func (c *Client) Restart() (*RestartResponse, error) {
	resp, err := c.Request("PUT", "/restart", &RequestOption{})
	if err != nil {
		return nil, err
	}

	var restartResp *RestartResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&restartResp)

	return restartResp, nil
}

// httpRequest ...
func (c *Client) httpRequest(method, path string, option *RequestOption) (*http.Response, error) {
	qs, err := query.Values(option.Query)
	if err != nil {
		return nil, err
	}

	host := fmt.Sprintf("%s:%d", c.Host, c.Port)
	u := &url.URL{Scheme: "http", Host: host, Path: c.BasePath + path, RawQuery: qs.Encode()}

	req, err := http.NewRequest(method, u.String(), nil) // todo body
	if err != nil {
		return nil, err
	}

	if c.UserAgent == "" {
		req.Header.Set("User-Agent", defaultUserAgent)
	} else {
		req.Header.Set("User-Agent", c.UserAgent+" "+defaultUserAgent)
	}

	if option.Priority == 0 {
		option.Priority = c.Priority
	}
	req.Header.Set("X-Mirakurun-Priority", strconv.Itoa(option.Priority))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// todo redirect

	return resp, nil
}

// requestStream ...
func (c *Client) requestStream(method, path string, option *RequestOption) (*http.Response, error) {
	resp, err := c.httpRequest(method, path, option)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 && resp.StatusCode > 202 {
		return nil, nil // todo error
	}

	return resp, nil
}

// getTS ...
func (c *Client) getTS(path string, decode bool) (*http.Response, error) {
	rawDecode := 0
	if decode {
		rawDecode = 1
	}
	option := &RequestOption{Query: &DecodableQuery{Decode: rawDecode}}

	resp, err := c.requestStream("GET", path, option)
	if err != nil {
		return nil, err
	}

	if resp.Header.Get("Content-Type") != "video/MP2T" {
		return nil, nil // todo error
	}

	return resp, nil
}
