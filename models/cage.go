package models

import "time"

type Cage struct {
	Id        int64     `json:"id" pg:"id,pk"`
	Name      string    `json:"name" pg:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt" pg:"created_at"`
}
