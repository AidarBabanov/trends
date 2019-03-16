package models

import "time"

type Trends struct {
	Topic string
	TrackedAt *time.Time
	Value int64
}