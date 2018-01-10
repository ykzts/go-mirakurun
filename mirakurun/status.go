package mirakurun

import (
	"context"
	"net/http"
)

// StatusService ...
type StatusService service

// Status ...
type Status struct {
	Version       string              `json:"version"`
	Process       ProcessStatus       `json:"process"`
	EPG           EPGStatus           `json:"epg"`
	StreamCount   map[string]int      `json:"streamCount"`
	ErrorCount    map[string]int      `json:"errorCount"`
	TimerAccuracy TimerAccuracyStatus `json:"timerAccuracy"`
}

// ProcessStatus ...
type ProcessStatus struct {
	Arch        string            `json:"arch"`
	Platform    string            `json:"platform"`
	Versions    map[string]string `json:"versions"`
	PID         int               `json:"pid"`
	MemoryUsage map[string]int    `json:"memoryUsage"`
}

// EPGStatus ...
type EPGStatus struct {
	GatheringNetworks []int `json:"gatheringNetworks"`
	StoredEvents      int   `json:"storedEvents"`
}

// TimerAccuracyStatus ...
type TimerAccuracyStatus struct {
	Last float64            `json:"last"`
	M1   map[string]float64 `json:"m1"`
	M5   map[string]float64 `json:"m5"`
	M15  map[string]float64 `json:"m15"`
}

// Get ...
func (s *StatusService) Get(ctx context.Context) (*Status, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "status", nil)
	if err != nil {
		return nil, nil, err
	}

	status := new(Status)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}