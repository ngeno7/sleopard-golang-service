package models

import "time"

type OutboundMessage struct {
	Id int64 `json:"id"`
	CampaignId int64 `json:"campaign_id"`
	CustomerId int64 `json:"customer_id"`
	Status string `json:"status"`
	RenderedContext string `json:"rendered_content"`
	LastError string `json:"last_error"`
	RetryCount int `json:"retry_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
