package models

import (
	"time"
)


type Campaign struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Channel string `json:"channel"`
	Status string `json:"status,omitempty"`
	BaseTemplate string `json:"base_template"`
	ScheduledAt *time.Time `json:"scheduled_at"`
	CreatedAt time.Time `json:"created_at"`
}