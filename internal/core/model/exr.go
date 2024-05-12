package model

import "time"

type Exr struct {
	ID           int       `json:"id"`
	CurCode      string    `json:"curCode"`
	DealBaseRate float64   `json:"dealBaseRate"`
	IsVirtual    bool      `json:"isVirtual"`
	CreatedAt    time.Time `json:"createdAt"`

	Subscribes    []Subscribe
	Notifications []Notification
}
