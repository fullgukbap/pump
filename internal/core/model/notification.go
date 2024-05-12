package model

import "time"

type Notification struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	UserID int `json:"userID"`
	ExrID  int `json:"exrID"`
}
