package models

import "time"

type Trends struct {
	Topic     string `json:"name"`
	TrackedAt *time.Time
	Value     int64 `json:"value"`
}
