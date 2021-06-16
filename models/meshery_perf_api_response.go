package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// PerformanceProfilesAPIResponse response retruned by performance endpoint on meshery server
type PerformanceProfilesAPIResponse struct {
	Page       uint                 `json:"page"`
	PageSize   uint                 `json:"page_size"`
	TotalCount uint                 `json:"total_count"`
	Profiles   []PerformanceProfile `json:"profiles,omitempty"`
}

// PerformanceResultsAPIResponse response retruned by performance endpoint on meshery server
type PerformanceResultsAPIResponse struct {
	Page       uint                `json:"page"`
	PageSize   uint                `json:"page_size"`
	TotalCount uint                `json:"total_count"`
	Results    []PerformanceResult `json:"results,omitempty"`
}

// PerformanceResult represents the result of a performance test
type PerformanceResult struct {
	MesheryID          *uuid.UUID    `json:"meshery_id,omitempty"`
	Name               string        `json:"name,omitempty"`
	Mesh               string        `json:"mesh,omitempty"`
	PerformanceProfile *uuid.UUID    `json:"performance_profile,omitempty"`
	UserID             *uuid.UUID    `json:"user_id"`
	RunnerResults      RunnerResults `json:"runner_results"`
	ServerMatrics      interface{}   `json:"server_metrics"`
	ServerBoardConfig  interface{}   `json:"server_board_config,omitempty"`
	TestStartTime      *time.Time    `json:"test_start_time,omitempty"`
}

type RunnerResults struct {
	URL               string     `json:"URL"`
	LoadGenerator     string     `json:"load-generator"`
	Duration          uint64     `json:"ActualDuration"`
	QPS               float64    `json:"ActualQPS"`
	StartTime         *time.Time `json:"StartTime"`
	DurationHistogram struct {
		Average     float64 `json:"Avg,omitempty"`
		Max         float64 `json:"Max,omitempty"`
		Min         float64 `json:"Min,omitempty"`
		Percentiles []struct {
			Percentile float64 `json:"Percentile,omitempty"`
			Value      float64 `json:"Value,omitempty"`
		} `json:"Percentiles,omitempty"`
	} `json:"DurationHistogram,omitempty"`
}
