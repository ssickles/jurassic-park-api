package models

import "time"

type PowerStatus string

const (
	PowerStatusActive PowerStatus = "ACTIVE"
	PowerStatusDown   PowerStatus = "DOWN"
)

type Cage struct {
	Id          int64       `json:"id" pg:"id,pk"`
	Name        string      `json:"name" pg:"name"`
	Capacity    int         `json:"capacity" pg:"capacity"`
	PowerStatus PowerStatus `json:"powerStatus" pg:"power_status"`
	CreatedAt   time.Time   `json:"createdAt" pg:"created_at"`
}
